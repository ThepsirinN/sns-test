package repositories

import (
	"context"
	"errors"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"

	"gorm.io/gorm/clause"
)

func (r *repoV1) UpdateLike(ctx context.Context, model models.Post, like entities.Like) error {
	var modelData models.Post
	modelData.Id = model.Id
	tx := r.db.Clauses(clause.Locking{Strength: "UPDATE"}).Limit(1).Find(&modelData)
	if err := tx.Error; err != nil {
		return err
	}
	for i := range modelData.Like {
		if modelData.Like[i].UserId == like.UserId {
			return errors.New("already like this post")
		}
	}

	modelData.OwnerId = like.UserId
	modelData.Like = append(modelData.Like, like)
	return tx.Save(&modelData).Error
}

func (r *repoV1) DeleteLike(ctx context.Context, model models.Post, like entities.Like) error {
	var modelData models.Post
	modelData.Id = model.Id
	tx := r.db.Clauses(clause.Locking{Strength: "UPDATE"}).Limit(1).Find(&modelData)
	if err := tx.Error; err != nil {
		return err
	}

	var newLike []entities.Like

	for i := range modelData.Like {
		if modelData.Like[i].UserId == like.UserId {
			continue
		}
		newLike = append(newLike, entities.Like{
			Id:     modelData.Like[i].Id,
			UserId: modelData.Like[i].UserId,
		})
	}

	modelData.OwnerId = like.UserId
	modelData.Like = newLike
	return tx.Save(&modelData).Error
}
