package models

import "time"

// can join this model with user model

type Friend struct {
	Id        int32      `gorm:"column:id;primarykey;autoIncrement"`
	SourceId  int32      `gorm:"column:source_id;NOT NULL;index"`
	DestId    int32      `gorm:"column:destination_id;NOT NULL;index"`
	Status    int        `gorm:"column:status;type:int(1);NOT NULL"` // 1:pending , 2:success
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP()"`
}

func (Friend) TableName() string {
	return "friend"
}
