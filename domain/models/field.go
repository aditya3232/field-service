package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Field struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	UUID          uuid.UUID      `gorm:"type:uuid;not null"`
	Code          string         `gorm:"type:varchar(15);not null"`
	Name          string         `gorm:"type:varchar(100);not null"`
	PricePerHour  int            `gorm:"type:int;not null"`
	Images        pq.StringArray `gorm:"type:text[]; not null"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     *gorm.DeletedAt
	FieldSchedule []FieldSchedule `gorm:"foreignKey:field_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
