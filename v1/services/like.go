package services

import (
	"context"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"

	"github.com/google/uuid"
)

func (s *serviceV1) AddLike(ctx context.Context, req entities.AddLikeRequest) error {
	uuidData, _ := uuid.NewV7()
	model := models.Post{
		Id: req.PostId,
	}

	likeData := entities.Like{
		Id:     uuidData.String(),
		UserId: req.UserId,
	}

	err := s.repoV1.UpdateLike(ctx, model, likeData)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) DeleteLike(ctx context.Context, req entities.DeleteLikeRequest) error {
	model := models.Post{
		Id: req.PostId,
	}
	likeData := entities.Like{
		UserId: req.UserId,
	}

	err := s.repoV1.DeleteLike(ctx, model, likeData)
	if err != nil {
		return err
	}

	return nil
}
