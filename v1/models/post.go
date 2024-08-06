package models

import (
	"sns-barko/v1/entities"
	"time"
)

type Post struct {
	Id        int32              `gorm:"column:id;primarykey;autoIncrement"`
	OwnerId   int32              `gorm:"column:owner_id;type:int;NOT NULL;index"`
	PostData  string             `gorm:"column:post_data;NOT NULL;type:varchar(255)"`
	PostImg   *string            `gorm:"column:post_img;type:varchar(255)"`
	Comment   []entities.Comment `gorm:"column:comment;serializer:json"`
	Like      []entities.Like    `gorm:"column:like;serializer:json"`
	CreatedAt *time.Time         `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP()"`
	UpdatedAt *time.Time         `gorm:"column:updated_at;type:timestamp;default:null"`
}

func (Post) TableName() string {
	return "post"
}
