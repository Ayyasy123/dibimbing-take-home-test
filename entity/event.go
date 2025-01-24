package entity

import "time"

type Event struct {
	ID                 int       `json:"id" gorm:"primary_key,auto_increment" `
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Location           string    `json:"location"`
	Date               time.Time `json:"date"`
	Category           string    `json:"category"`
	Capacity           int       `json:"capacity"`
	Price              int       `json:"price"`
	Status             string    `json:"status" gorm:"default:active"`
	AvailableTickets   int       `json:"available_tickets"`
	TicketAvailability string    `json:"ticket_availability"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Tickets            []Ticket  `json:"tickets,omitempty" gorm:"foreignKey:EventID"`
}

type CreateEventReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Date        string `json:"date" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Capacity    int    `json:"capacity" validate:"required,gte=0"`
	Price       int    `json:"price" validate:"required,gte=0"`
}

type UpdateEventReq struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Category           string `json:"category"`
	Capacity           int    `json:"capacity" validate:"gte=0"`
	Price              int    `json:"price" validate:"gte=0"`
	Status             string `json:"status"`
	AvailableTickets   int    `json:"available_tickets"`
	TicketAvailability string `json:"ticket_availability"`
}

type EventRes struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Category           string `json:"category"`
	Capacity           int    `json:"capacity"`
	Price              int    `json:"price"`
	Status             string `json:"status"`
	AvailableTickets   int    `json:"available_tickets"`
	TicketAvailability string `json:"ticket_availability"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

type EventStatusDistribution struct {
	EventStatus   string `json:"event_status"`   // Status event (Aktif, Berlangsung, Selesai, Dibatalkan)
	TotalCapacity int    `json:"total_capacity"` // Total kapasitas event dengan status tersebut
	TicketBooked  int    `json:"ticket_booked"`  // Total tiket yang sudah dipesan
}

type EventReport struct {
	TotalEvent              int                       `json:"total_event"`               // Total event yang tersedia
	EventStatusDistribution []EventStatusDistribution `json:"event_status_distribution"` // Distribusi status event
}
