package entities

type AddLikeRequest struct {
	Id     string
	PostId int32 `json:"post_id" validate:"required"`
	UserId int32 `validate:"required"`
}

type DeleteLikeRequest struct {
	PostId int32 `json:"post_id" validate:"required"`
	UserId int32 `validate:"required"`
}
