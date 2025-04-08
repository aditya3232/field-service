package models

import (
	"github.com/google/uuid"
	"time"
)

type Time struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;not null"`
	StartTime string    `gorm:"type:time without time zone;not null"`
	EndTime   string    `gorm:"type:time without time zone;not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
