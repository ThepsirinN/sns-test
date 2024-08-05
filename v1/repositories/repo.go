package repositories

import (
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
