package models

import (
	"time"
)

type User struct {
	Id         int32      `gorm:"column:id;primarykey;autoIncrement"`
	Email      string     `gorm:"column:email;type:varchar(30);NOT NULL;uniqueIndex"`
	Firstname  string     `gorm:"column:first_name;type:varchar(60);NOT NULL"`
	Lastname   string     `gorm:"column:last_name;type:varchar(60);NOT NULL"`
	ProfileImg *string    `gorm:"column:img_profile"`
	Auth       string     `gorm:"column:auth;type:varchar(250);NOT NULL"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP()"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;type:timestamp;default:null"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:timestamp;index"`
}

func (User) TableName() string {
	return "user"
}
