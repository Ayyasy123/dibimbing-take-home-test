package repository

import (
	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *entity.Event) error
	FindEventByID(id int) (*entity.Event, error)
	FindAllEvents() ([]entity.Event, error)
	UpdateEvent(id int, event *entity.Event) error
	DeleteEvent(id int) error
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
