package entities

type CreateUserRequest struct {
	Email       string  `json:"email" validate:"required,max=30,email"`
	Firstname   string  `json:"first_name" validate:"required,max=60"`
	Lastname    string  `json:"last_name" validate:"required,max=60"`
	ProfileImg  *string `json:"profile_image"`
	Auth        string  `json:"password" validate:"required,min=8,max=20"`
	ConfirmPass string  `json:"confirm_password" validate:"required,min=8,max=20"`
}

type (
	AuthUserRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	AuthUserResponse struct {
		JWT string `json:"user_data"`
	}
)
type (
	FindUserByEmailRequest struct {
		Email string `validate:"required"`
	}

	FindUserByEmailResponse struct {
		Id         int32   `json:"id"`
		Email      string  `json:"email"`
		Firstname  string  `json:"first_name"`
		Lastname   string  `json:"last_name"`
		ProfileImg *string `json:"profile_img"`
	}
)

type (
	UpdateUserRequest struct {
		Id          int32   `validate:"required"`
		Firstname   string  `json:"first_name" validate:"max=60"`
		Lastname    string  `json:"last_name" validate:"max=60"`
		ProfileImg  *string `json:"profile_image"`
		Auth        *string `json:"password"`
		ConfirmPass *string `json:"confirm_password"`
	}

	UpdateUserResponse struct {
		JWT string `json:"user_data"`
	}
)
type DeleteUserRequest struct {
	Id int32 `validte:"required"`
}
