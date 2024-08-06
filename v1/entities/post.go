package entities

import "time"

type CreatePostRequest struct {
	OwnerId  int32   `validate:"required"`
	PostData string  `json:"post_data" validate:"required"`
	PostImg  *string `json:"post_img"`
}

type (
	ListAllPostFromUserRequest struct {
		OwnerId int32 `validate:"required"`
	}

	ListAllPostFromUserResponse struct {
		Id              int32      `json:"id"`
		OwnerId         int32      `json:"owner_id"`
		OwnerFirstName  string     `json:"owner_first_name"`
		OwnerLastName   string     `json:"owner_last_name"`
		OwnerProfileImg *string    `json:"owner_profile_img"`
		PostData        string     `json:"post_data"`
		PostImg         *string    `json:"post_img"`
		Comment         []Comment  `json:"comment" gorm:"serializer:json"`
		Like            []Like     `json:"like" gorm:"serializer:json"`
		CreatedAt       *time.Time `json:"created_at"`
		UpdatedAt       *time.Time `json:"updated_at"`
	}
)

type (
	ReadPostByPostIdRequest struct {
		Id      int32 `param:"id" validate:"required"`
		OwnerId int32 `validate:"required"`
	}

	ReadPostByPostIdResponse struct {
		Id              int32      `json:"id"`
		OwnerId         int32      `json:"owner_id"`
		OwnerFirstName  string     `json:"owner_first_name"`
		OwnerLastName   string     `json:"owner_last_name"`
		OwnerProfileImg *string    `json:"owner_profile_img"`
		PostData        string     `json:"post_data"`
		PostImg         *string    `json:"post_img"`
		Comment         []Comment  `json:"comment" gorm:"serializer:json"`
		Like            []Like     `json:"like" gorm:"serializer:json"`
		CreatedAt       *time.Time `json:"created_at"`
		UpdatedAt       *time.Time `json:"updated_at"`
	}
)

type UpdatePostRequest struct {
	Id       int32   `json:"id" validate:"required"`
	OwnerId  int32   `validate:"required"`
	PostData string  `json:"post_data"`
	PostImg  *string `json:"post_img"`
}

type DeletePostRequest struct {
	Id      int32 `json:"id" validate:"required"`
	OwnerId int32 `validate:"required"`
}

type Comment struct {
	Id            string    `json:"id"`
	UserId        int32     `json:"user_id"`
	UserFirstName string    `json:"user_first_name"`
	UserLastName  string    `json:"user_last_name"`
	UserImg       *string   `json:"user_img_profile"`
	CommentData   string    `json:"comment_data"`
	CommentImg    *string   `json:"comment_img"`
	CreateAt      time.Time `json:"create_at"`
}

type Like struct {
	Id     string `json:"id"`
	UserId int32  `json:"user_id"`
}
