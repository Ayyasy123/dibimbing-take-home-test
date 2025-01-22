package entity

import "time"

type Ticket struct {
	ID        int       `json:"id" gorm:"primary_key,auto_increment" `
	EventID   int       `json:"event_id" gorm:"not null" `
	UserID    int       `json:"user_id" gorm:"not null" `
	Status    string    `json:"status" gorm:"default:Dibeli" `
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
	Event     Event     `json:"event" gorm:"foreignKey:EventID" `
	User      User      `json:"user" gorm:"foreignKey:UserID" `
}

type CreateTicketReq struct {
	EventID int    `json:"event_id" validate:"required"`
	UserID  int    `json:"user_id" validate:"required"`
	Status  string `json:"status" validate:"required"`
}

type UpdateTicketReq struct {
	Status string `json:"status" validate:"required"`
}

type TicketRes struct {
	ID        int    `json:"id"`
	EventID   int    `json:"event_id"`
	UserID    int    `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
