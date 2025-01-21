package service

import (
	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/repository"
)

type EventService interface {
	CreateEvent(req *entity.CreateEventReq) (*entity.Event, error)
	FindEventByID(id int) (*entity.Event, error)
	FindAllEvents() ([]entity.Event, error)
	UpdateEvent(req *entity.UpdateEventReq) error
	DeleteEvent(id int) error
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventService{eventRepository: eventRepository}
}

func (s *eventService) CreateEvent(req *entity.CreateEventReq) (*entity.Event, error) {
	event := &entity.Event{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Capacity:    req.Capacity,
		Price:       req.Price,
		Status:      "active",
	}

	err := s.eventRepository.CreateEvent(event)
	return event, err
}

func (s *eventService) FindEventByID(id int) (*entity.Event, error) {
	return s.eventRepository.FindEventByID(id)
}

func (s *eventService) FindAllEvents() ([]entity.Event, error) {
	return s.eventRepository.FindAllEvents()
}

func (s *eventService) UpdateEvent(req *entity.UpdateEventReq) error {

	service, err := s.eventRepository.FindEventByID(req.ID)
	if err != nil {
		return err
	}
	service.Name = req.Name
	service.Description = req.Description
	service.Location = req.Location
	service.Capacity = req.Capacity
	service.Price = req.Price

	return s.eventRepository.UpdateEvent(service)

}

func (s *eventService) DeleteEvent(id int) error {
	return s.eventRepository.DeleteEvent(id)
}
