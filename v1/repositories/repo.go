package repositories

import (
	"context"
	"sns-barko/v1/models"
	"time"

	"gorm.io/gorm"
)

type repoV1 struct {
	db  *gorm.DB
	now time.Time
}

func New(db *gorm.DB) *repoV1 {
	repo := repoV1{db: db, now: time.Now().Local()}
	return &repo
}

func (r *repoV1) AutoMigrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&models.User{}, &models.Friend{}, &models.Post{})
}
