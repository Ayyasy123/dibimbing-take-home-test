package entity

import "time"

type Event struct {
	ID          int       `json:"id" gorm:"primary_key,auto_increment" `
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
	Price       int       `json:"price"`
	Status      string    `json:"status" gorm:"default:active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tickets     []Ticket  `json:"tickets" gorm:"foreignKey:EventID"`
}

type CreateEventReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Capacity    int    `json:"capacity" validate:"required,gte=0"`
	Price       int    `json:"price" validate:"required,gte=0"`
}

type UpdateEventReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Capacity    int    `json:"capacity" validate:"required,gte=0"`
	Price       int    `json:"price" validate:"required,gte=0"`
}

type EventRes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Capacity    int    `json:"capacity"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
