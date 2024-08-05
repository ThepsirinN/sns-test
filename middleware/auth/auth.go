package auth

import (
	"errors"
	"fmt"
	"sns-barko/utility/logger"
	"sns-barko/utility/response"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserClaim struct {
	jwt.RegisteredClaims
	Id         int32   `json:"id"`
	Firstname  string  `json:"first_name"`
	Lastname   string  `json:"last_name"`
	ProfileImg *string `json:"profile_image"`
}

func MiddleWareAuth(JWTsecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			authorizationToken := req.Header.Get("Authorization")
			errCode, errResp := response.NewUnAuthorizeResponse(fmt.Errorf("found an error on auth process").Error())

			if authorizationToken == "" {
				logger.Error(ctx, errors.New("authorizationToken is empty"))
				return c.JSON(errCode, errResp)
			}

			token, err := jwt.ParseWithClaims(
				authorizationToken,
				&UserClaim{},
				func(token *jwt.Token) (interface{}, error) {
					return []byte(JWTsecret), nil
				},
				jwt.WithLeeway(5*time.Second),
				jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
				jwt.WithExpirationRequired(),
			)

			if err != nil {
				logger.Error(ctx, err)
				return c.JSON(errCode, errResp)
			}

			if token != nil {
				if !token.Valid {
					logger.Error(ctx, errors.New("invalid token"))
					return c.JSON(errCode, errResp)
				}
			}

			var userClaim *UserClaim
			unAuthoriazeCode, unAuthoriazeResp := response.NewUnAuthorizeWithOutDataResponse()
			claim, ok := token.Claims.(*UserClaim)
			if !ok {
				logger.Error(ctx, errors.New("invalid token structure"))
				return c.JSON(unAuthoriazeCode, unAuthoriazeResp)
			}

			if time.Now().Local().After(claim.ExpiresAt.Time) {
				logger.Error(ctx, errors.New("token expire"))
				return c.JSON(unAuthoriazeCode, unAuthoriazeResp)
			}

			userClaim = claim

			c.Set("jwt_user", userClaim)
			c.Set("user_id", userClaim.Id)

			// next
			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}
