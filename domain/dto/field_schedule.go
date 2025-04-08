package dto

import (
	"field-service/constants"
	"github.com/google/uuid"
	"time"
)

type FieldScheduleRequest struct {
	FieldID string   `json:"fieldID" validate:"required"`
	Date    string   `json:"date" validate:"required"`
	TimeIDs []string `json:"timeIDs" validate:"required"`
}

type GenerateFieldScheduleForOneMonthRequest struct {
	FieldID string `json:"fieldID" validate:"required"`
}

type UpdateFieldScheduleRequest struct {
	Date   string `json:"date" validate:"required"`
	TimeID string `json:"timeID" validate:"required"`
}

type UpdateStatusFieldScheduleRequest struct {
	FieldScheduleIDs []string `json:"fieldScheduleIDs" validate:"required"`
}

type FieldScheduleResponse struct {
	UUID         uuid.UUID                         `json:"uuid"`
	FieldName    string                            `json:"fieldName"`
	PricePerHour int                               `json:"pricePerHour"`
	Date         string                            `json:"date"`
	Status       constants.FieldScheduleStatusName `json:"status"`
	Time         string                            `json:"time"`
	CreatedAt    *time.Time                        `json:"createdAt"`
	UpdatedAt    *time.Time                        `json:"updatedAt"`
}

type FieldScheduleForBookingResponse struct {
	UUID         uuid.UUID                         `json:"uuid"`
	PricePerHour string                            `json:"pricePerHour"`
	Date         string                            `json:"date"`
	Status       constants.FieldScheduleStatusName `json:"status"`
	Time         string                            `json:"time"`
}

type FieldScheduleRequestParam struct {
	Page       int     `form:"page" validate:"required"`
	Limit      int     `form:"limit" validate:"required"`
	SortColumn *string `form:"sortColumn"`
	SortOrder  *string `form:"sortOrder"`
}

type FieldScheduleByFieldIDAndDateRequestParam struct {
	Date string `form:"date" validate:"required"`
}
