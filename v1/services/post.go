package services

import (
	"context"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"
)

func (s *serviceV1) CreatePost(ctx context.Context, req entities.CreatePostRequest) error {
	model := models.Post{
		OwnerId:  req.OwnerId,
		PostData: req.PostData,
		PostImg:  req.PostImg,
	}

	err := s.repoV1.CreatePost(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) ListAllPostFromUser(ctx context.Context, req entities.ListAllPostFromUserRequest, resp *[]entities.ListAllPostFromUserResponse) error {
	err := s.repoV1.ListAllPostFromUser(ctx, req.OwnerId, resp)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceV1) ReadPostByPostId(ctx context.Context, req entities.ReadPostByPostIdRequest, resp *entities.ReadPostByPostIdResponse) error {
	resp.Id = req.Id
	resp.OwnerId = req.OwnerId

	err := s.repoV1.ReadPostByPostId(ctx, resp)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) UpdatePostData(ctx context.Context, req entities.UpdatePostRequest) error {
	model := models.Post{
		Id:       req.Id,
		OwnerId:  req.OwnerId,
		PostData: req.PostData,
		PostImg:  req.PostImg,
	}

	err := s.repoV1.UpdatePostData(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) DeletePost(ctx context.Context, req entities.DeletePostRequest) error {
	model := models.Post{
		Id:      req.Id,
		OwnerId: req.OwnerId,
	}

	err := s.repoV1.DeletePost(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
