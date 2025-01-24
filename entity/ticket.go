package entity

import (
	"database/sql"
	"time"
)

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

type TicketStatusDistribution struct {
	Status       string `json:"status"`        // Status tiket (Dibeli, Dibatalkan)
	TotalTickets int    `json:"total_tickets"` // Total tiket dengan status tersebut
	TotalRevenue int    `json:"total_revenue"` // Total pendapatan dari tiket dengan status tersebut
}

type TicketReport struct {
	TotalTickets             int                        `json:"total_tickets"`              // Total tiket yang terjual
	TotalRevenue             int                        `json:"total_revenue"`              // Total pendapatan dari tiket yang terjual
	TicketStatusDistribution []TicketStatusDistribution `json:"ticket_status_distribution"` // Distribusi status tiket
}

type TicketStatusDistributionResult struct {
	TotalTickets sql.NullInt64 `gorm:"column:total_tickets" json:"total_tickets`
	TotalRevenue sql.NullInt64 `gorm:"column:total_revenue" json:"total_revenue"`
}

type TicketsSoldPerEvent struct {
	EventID      int    `json:"event_id"`
	EventName    string `json:"event_name"`
	TotalTickets int    `json:"total_tickets"`
	TotalRevenue int    `json:"total_revenue"`
}
