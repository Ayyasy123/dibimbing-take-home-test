package service

import (
	"errors"
	"time"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/repository"
)

type EventService interface {
	CreateEvent(req *entity.CreateEventReq) (*entity.Event, error)
	FindEventByID(id int) (*entity.Event, error)
	FindAllEvents() ([]entity.Event, error)
	UpdateEvent(id int, req *entity.UpdateEventReq) error
	DeleteEvent(id int) error
	SearchEvents(searchQuery string, minPrice, maxPrice int, category, status string, startDate, endDate time.Time) ([]entity.EventRes, error)
	GetEventReport(startDate, endDate time.Time) (*entity.EventReport, error)
	CancelEvent(eventID int) error
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{eventRepository: eventRepository}
}

func (s *eventService) CreateEvent(req *entity.CreateEventReq) (*entity.Event, error) {
	// Parse date
	eventDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format, must be YYYY-MM-DD")
	}

	existingEventName, err := s.eventRepository.IsEventNameExists(req.Name)
	if err != nil {
		return nil, err
	}

	if existingEventName {
		return nil, errors.New("event name already exists")
	}

	event := &entity.Event{
		Name:               req.Name,
		Description:        req.Description,
		Location:           req.Location,
		Date:               eventDate,
		Category:           req.Category,
		Capacity:           req.Capacity,
		Price:              req.Price,
		Status:             "Aktif",
		AvailableTickets:   req.Capacity,
		TicketAvailability: "Tersedia",
	}

	err = s.eventRepository.CreateEvent(event)
	return event, err
}

func (s *eventService) FindEventByID(id int) (*entity.Event, error) {
	return s.eventRepository.FindEventByID(id)
}

func (s *eventService) FindAllEvents() ([]entity.Event, error) {
	return s.eventRepository.FindAllEvents()
}

func (s *eventService) UpdateEvent(id int, req *entity.UpdateEventReq) error {
	existingEvent, err := s.eventRepository.FindEventByID(id)
	if err != nil {
		return err
	}

	if req.Name != "" {
		existingEvent.Name = req.Name
	}

	if req.Description != "" {
		existingEvent.Description = req.Description
	}

	if req.Location != "" {
		existingEvent.Location = req.Location
	}

	if req.Category != "" {
		existingEvent.Category = req.Category
	}

	if req.Capacity != 0 {
		existingEvent.Capacity = req.Capacity
	}

	if req.Price != 0 {
		existingEvent.Price = req.Price
	}

	if req.Status != "" {
		existingEvent.Status = req.Status
	}

	if req.AvailableTickets != 0 {
		existingEvent.AvailableTickets = req.AvailableTickets
	}

	if req.TicketAvailability != "" {
		existingEvent.TicketAvailability = req.TicketAvailability
	}

	return s.eventRepository.UpdateEvent(id, existingEvent)
}

func (s *eventService) DeleteEvent(id int) error {
	return s.eventRepository.DeleteEvent(id)
}

func (s *eventService) SearchEvents(searchQuery string, minPrice, maxPrice int, category, status string, startDate, endDate time.Time) ([]entity.EventRes, error) {
	events, err := s.eventRepository.SearchEvents(searchQuery, minPrice, maxPrice, category, status, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var eventRes []entity.EventRes
	for _, event := range events {
		eventRes = append(eventRes, entity.EventRes{
			ID:                 event.ID,
			Name:               event.Name,
			Description:        event.Description,
			Location:           event.Location,
			Date:               event.Date.Format("2006-01-02 15:04:05"),
			Category:           event.Category,
			Capacity:           event.Capacity,
			Price:              event.Price,
			Status:             event.Status,
			AvailableTickets:   event.AvailableTickets,
			TicketAvailability: event.TicketAvailability,
			CreatedAt:          event.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:          event.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return eventRes, nil
}

func (s *eventService) GetEventReport(startDate, endDate time.Time) (*entity.EventReport, error) {
	// Hitung total event
	totalEvent, err := s.eventRepository.GetTotalEvents(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Daftar status event yang ingin dihitung
	statuses := []string{"Aktif", "Berlangsung", "Selesai", "Dibatalkan"}

	// Slice untuk menyimpan distribusi status event
	var statusDistribution []entity.EventStatusDistribution

	// Loop melalui setiap status dan hitung distribusinya
	for _, status := range statuses {
		distribution, err := s.eventRepository.GetEventStatusDistribution(status, startDate, endDate)
		if err != nil {
			return nil, err
		}
		distribution.EventStatus = status
		statusDistribution = append(statusDistribution, distribution)
	}

	return &entity.EventReport{
		TotalEvent:              int(totalEvent),
		EventStatusDistribution: statusDistribution,
	}, nil
}

func (s *eventService) CancelEvent(eventID int) error {
	// Cek apakah event ada
	event, err := s.eventRepository.FindEventByID(eventID)
	if err != nil {
		return err
	}

	// Validasi: Event hanya bisa dibatalkan jika statusnya "active" atau "upcoming"
	if event.Status != "active" && event.Status != "upcoming" {
		return errors.New("event cannot be cancelled because it is not in 'active' or 'upcoming' status")
	}

	// Batalkan event dan semua tiket terkait
	return s.eventRepository.CancelEvent(eventID)
}
