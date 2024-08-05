package services

import (
	"context"
	"errors"
	"fmt"
	"sns-barko/utility/ptr"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *serviceV1) CreateUser(ctx context.Context, req entities.CreateUserRequest) error {

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(req.Auth+s.secret.User.Hash.Secret), s.secret.User.Hash.Cost)
	if err != nil {
		return err
	}

	model := models.User{
		Email:      req.Email,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		ProfileImg: req.ProfileImg,
		Auth:       string(hashBytes),
	}

	err = s.userRepo.CreateUser(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) AuthUser(ctx context.Context, req entities.AuthUserRequest, resp *entities.AuthUserResponse) error {
	model := models.User{
		Email: req.Email,
	}

	err := s.userRepo.ReadUsersByEmail(ctx, &model)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Auth), []byte(fmt.Sprint(req.Password, s.secret.User.Hash.Secret)))
	if err != nil {
		return errors.New("email and password miss match")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":            model.Id,
			"first_name":    model.Firstname,
			"last_name":     model.Lastname,
			"profile_image": model.ProfileImg,
			"exp":           (time.Now().Add(s.secret.User.JWT.ExpDuration)).Unix(),
		})

	tokenString, err := token.SignedString([]byte(s.secret.User.JWT.Secret))
	if err != nil {
		return err
	}

	resp.JWT = tokenString

	return nil
}

func (s *serviceV1) FindUsersByEmail(ctx context.Context, userId int32, req entities.FindUserByEmailRequest, resp *[]entities.FindUserByEmailResponse) error {
	var model []models.User
	err := s.userRepo.FindUsersByEmail(ctx, req.Email, userId, &model)
	if err != nil {
		return err
	}
	i, length := 0, len(model)
	for i < length {
		*resp = append(*resp, entities.FindUserByEmailResponse{
			Email:      model[i].Email,
			Firstname:  model[i].Firstname,
			Lastname:   model[i].Lastname,
			ProfileImg: model[i].ProfileImg,
		})
		i++
	}
	return nil
}

func (s *serviceV1) UpdateUser(ctx context.Context, req entities.UpdateUserRequest, resp *entities.UpdateUserResponse) error {

	model := models.User{
		Id:         req.Id,
		Firstname:  req.Firstname,
		Lastname:   req.Lastname,
		ProfileImg: req.ProfileImg,
		Auth:       "",
	}

	if req.Auth != nil && req.ConfirmPass != nil {
		if len(*req.Auth) < 8 || len(*req.Auth) > 20 {
			return errors.New("password must less than 20 and more than 8 chars")
		}

		if *req.Auth != *req.ConfirmPass {
			return errors.New("password and confirm password miss match")
		}

		hashBytes, err := bcrypt.GenerateFromPassword([]byte(*req.Auth+s.secret.User.Hash.Secret), s.secret.User.Hash.Cost)
		if err != nil {
			return err
		}

		model.Auth = string(hashBytes)
	}

	err := s.userRepo.UpdateUsers(ctx, model)
	if err != nil {
		return err
	}

	err = s.userRepo.ReadUsersById(ctx, &model)

	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":            model.Id,
			"first_name":    model.Firstname,
			"last_name":     model.Lastname,
			"profile_image": model.ProfileImg,
			"exp":           (time.Now().Add(s.secret.User.JWT.ExpDuration)).Unix(),
		})

	tokenString, err := token.SignedString([]byte(s.secret.User.JWT.Secret))
	if err != nil {
		return err
	}

	resp.JWT = tokenString

	return nil
}

func (s *serviceV1) DeleteUser(ctx context.Context, req entities.DeleteUserRequest) error {
	model := models.User{
		Id:        req.Id,
		DeletedAt: ptr.ToPointer(time.Now().Local()),
	}
	err := s.userRepo.DeleteUserById(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
