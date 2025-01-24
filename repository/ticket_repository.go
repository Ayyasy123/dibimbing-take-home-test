package repository

import (
	"database/sql"
	"time"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"gorm.io/gorm"
)

type TicketRepository interface {
	CreateTicket(ticket *entity.Ticket) error
	FindTicketByID(id int) (*entity.Ticket, error)
	FindAllTickets() ([]entity.Ticket, error)
	UpdateTicket(id int, ticket *entity.Ticket) error
	DeleteTicket(id int) error
	FindAllTicketsByUserID(userID int) ([]entity.Ticket, error)
	GetTotalTickets(startDate, endDate time.Time) (int64, error)
	GetTotalRevenue(startDate, endDate time.Time) (int, error)
	GetTicketStatusDistribution(status string, startDate, endDate time.Time) (int, int, error)
	GetTicketsSoldPerEvent(startDate, endDate time.Time, eventID int) ([]entity.TicketsSoldPerEvent, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *ticketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) CreateTicket(ticket *entity.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) FindTicketByID(id int) (*entity.Ticket, error) {
	var ticket entity.Ticket
	err := r.db.Where("id = ?", id).First(&ticket).Error
	return &ticket, err
}

func (r *ticketRepository) FindAllTickets() ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := r.db.Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) UpdateTicket(id int, ticket *entity.Ticket) error {
	result := r.db.Model(&entity.Ticket{}).Where("id = ?", id).Updates(ticket)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ticketRepository) DeleteTicket(id int) error {
	return r.db.Delete(&entity.Ticket{}, id).Error
}

func (r *ticketRepository) FindAllTicketsByUserID(userID int) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := r.db.Where("user_id = ?", userID).Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) GetTotalTickets(startDate, endDate time.Time) (int64, error) {
	var totalTickets int64
	query := r.db.Model(&entity.Ticket{})

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("created_at <= ?", endDate)
	}

	err := query.Count(&totalTickets).Error
	return totalTickets, err
}

func (r *ticketRepository) GetTotalRevenue(startDate, endDate time.Time) (int, error) {
	var totalRevenue sql.NullInt64 // Gunakan sql.NullInt64 untuk menangani NULL
	query := r.db.Model(&entity.Ticket{}).
		Joins("JOIN events ON tickets.event_id = events.id")

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("tickets.created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("tickets.created_at <= ?", endDate)
	}

	err := query.Select("SUM(events.price)").Scan(&totalRevenue).Error
	if err != nil {
		return 0, err
	}

	// Jika totalRevenue adalah NULL, kembalikan 0
	if !totalRevenue.Valid {
		return 0, nil
	}

	return int(totalRevenue.Int64), nil
}

func (r *ticketRepository) GetTicketStatusDistribution(status string, startDate, endDate time.Time) (int, int, error) {
	var result entity.TicketStatusDistributionResult
	query := r.db.Model(&entity.Ticket{}).
		Joins("JOIN events ON tickets.event_id = events.id").
		Where("tickets.status = ?", status)

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("tickets.created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("tickets.created_at <= ?", endDate)
	}

	err := query.Select("COUNT(tickets.id) as total_tickets, SUM(events.price) as total_revenue").
		Scan(&result).Error

	// Konversi sql.NullInt64 ke int
	totalTickets := 0
	if result.TotalTickets.Valid {
		totalTickets = int(result.TotalTickets.Int64)
	}

	totalRevenue := 0
	if result.TotalRevenue.Valid {
		totalRevenue = int(result.TotalRevenue.Int64)
	}

	return totalTickets, totalRevenue, err
}

func (r *ticketRepository) GetTicketsSoldPerEvent(startDate, endDate time.Time, eventID int) ([]entity.TicketsSoldPerEvent, error) {
	var results []entity.TicketsSoldPerEvent
	query := r.db.Model(&entity.Ticket{}).
		Joins("JOIN events ON tickets.event_id = events.id").
		Where("tickets.status = ?", "Dibeli"). // Hanya tiket dengan status "Dibeli"
		Group("events.id")

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("tickets.created_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("tickets.created_at <= ?", endDate)
	}

	// Tambahkan filter event_id jika tidak nol
	if eventID != 0 {
		query = query.Where("events.id = ?", eventID)
	}

	err := query.Select("events.id as event_id, events.name as event_name, COUNT(tickets.id) as total_tickets, SUM(events.price) as total_revenue").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
