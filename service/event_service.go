package service

import (
	"errors"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/repository"
)

type EventService interface {
	CreateEvent(req *entity.CreateEventReq) (*entity.Event, error)
	FindEventByID(id int) (*entity.Event, error)
	FindAllEvents() ([]entity.Event, error)
	UpdateEvent(id int, req *entity.UpdateEventReq) error
	DeleteEvent(id int) error
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{eventRepository: eventRepository}
}

func (s *eventService) CreateEvent(req *entity.CreateEventReq) (*entity.Event, error) {
	existingEventName, err := s.eventRepository.IsEventNameExists(req.Name)
	if err != nil {
		return nil, err
	}

	if existingEventName {
		return nil, errors.New("event name already exists")
	}

	event := &entity.Event{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Capacity:    req.Capacity,
		Price:       req.Price,
		Status:      "active",
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

	if req.Capacity != 0 {
		existingEvent.Capacity = req.Capacity
	}

	if req.Price != 0 {
		existingEvent.Price = req.Price
	}

	return s.eventRepository.UpdateEvent(id, existingEvent)
}

func (s *eventService) DeleteEvent(id int) error {
	return s.eventRepository.DeleteEvent(id)
}
