package repository

import (
	"strings"
	"time"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *entity.Event) error
	FindEventByID(id int) (*entity.Event, error)
	FindAllEvents() ([]entity.Event, error)
	UpdateEvent(id int, event *entity.Event) error
	DeleteEvent(id int) error
	IsEventNameExists(name string) (bool, error)
	SearchEvents(searchQuery string, minPrice, maxPrice int, category, status string, startDate, endDate time.Time) ([]entity.Event, error)
	GetTotalEvents(startDate, endDate time.Time) (int64, error)
	GetEventStatusDistribution(status string, startDate, endDate time.Time) (entity.EventStatusDistribution, error)
	CancelEvent(eventID int) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) CreateEvent(event *entity.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) FindEventByID(id int) (*entity.Event, error) {
	var event entity.Event
	err := r.db.Where("id = ?", id).First(&event).Error
	return &event, err
}

func (r *eventRepository) FindAllEvents() ([]entity.Event, error) {
	var events []entity.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventRepository) UpdateEvent(id int, event *entity.Event) error {
	result := r.db.Model(&entity.Event{}).Where("id = ?", id).Updates(event)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *eventRepository) DeleteEvent(id int) error {
	return r.db.Delete(&entity.Event{}, id).Error
}

func (r *eventRepository) IsEventNameExists(name string) (bool, error) {
	var count int64
	r.db.Model(&entity.Event{}).Where("name = ?", strings.ToLower(name)).Count(&count)
	return count > 0, nil
}

func (r *eventRepository) SearchEvents(searchQuery string, minPrice, maxPrice int, category, status string, startDate, endDate time.Time) ([]entity.Event, error) {
	var events []entity.Event

	query := r.db

	if searchQuery != "" {
		searchQuery = strings.ToLower(searchQuery)
		query = query.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	query = query.Where("price BETWEEN ? AND ?", minPrice, maxPrice)

	if category != "" {
		category = strings.ToLower(category)
		query = query.Where("LOWER(category) = ?", category)
	}

	if status != "" {
		status = strings.ToLower(status)
		query = query.Where("LOWER(status) = ?", status)
	}

	if !startDate.IsZero() {
		query = query.Where("date >= ?", startDate)
	}

	if !endDate.IsZero() {
		query = query.Where("date <= ?", endDate)
	}

	err := query.Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *eventRepository) GetTotalEvents(startDate, endDate time.Time) (int64, error) {
	var totalEvent int64
	query := r.db.Model(&entity.Event{})

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("date <= ?", endDate)
	}

	err := query.Count(&totalEvent).Error
	return totalEvent, err
}

func (r *eventRepository) GetEventStatusDistribution(status string, startDate, endDate time.Time) (entity.EventStatusDistribution, error) {
	var distribution entity.EventStatusDistribution
	query := r.db.Model(&entity.Event{}).Where("status = ?", status)

	// Tambahkan filter tanggal jika startDate atau endDate tidak kosong
	if !startDate.IsZero() {
		query = query.Where("date >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("date <= ?", endDate)
	}

	err := query.Select("SUM(capacity) as total_capacity, SUM(capacity - available_tickets) as ticket_booked").
		Scan(&distribution).Error
	return distribution, err
}

func (r *eventRepository) CancelEvent(eventID int) error {
	// Mulai transaksi database
	tx := r.db.Begin()

	// Update status event menjadi "cancelled"
	if err := tx.Model(&entity.Event{}).Where("id = ?", eventID).
		Update("status", "cancelled").Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update status semua tiket terkait menjadi "cancelled"
	if err := tx.Model(&entity.Ticket{}).Where("event_id = ?", eventID).
		Update("status", "cancelled").Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaksi
	return tx.Commit().Error
}
