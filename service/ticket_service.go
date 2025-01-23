package service

import (
	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/repository"
)

type TicketService interface {
	CreateTicket(req *entity.CreateTicketReq) (*entity.TicketRes, error)
	FindTicketByID(id int) (*entity.TicketRes, error)
	FindAllTickets() ([]entity.TicketRes, error)
	UpdateTicket(id int, req *entity.UpdateTicketReq) error
	DeleteTicket(id int) error
	FindAllTicketsByUserID(userID int) ([]entity.TicketRes, error)
}

type ticketService struct {
	ticketRepository repository.TicketRepository
}

func NewTicketService(ticketRepository repository.TicketRepository) TicketService {
	return &ticketService{ticketRepository: ticketRepository}
}

func (s *ticketService) CreateTicket(req *entity.CreateTicketReq) (*entity.TicketRes, error) {
	ticket := &entity.Ticket{
		EventID: req.EventID,
		UserID:  req.UserID,
		Status:  "Dibeli",
	}

	err := s.ticketRepository.CreateTicket(ticket)

	ticketRes := &entity.TicketRes{
		ID:        ticket.ID,
		EventID:   ticket.EventID,
		UserID:    ticket.UserID,
		Status:    ticket.Status,
		CreatedAt: ticket.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: ticket.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ticketRes, err
}

func (s *ticketService) FindTicketByID(id int) (*entity.TicketRes, error) {
	ticket, err := s.ticketRepository.FindTicketByID(id)
	if err != nil {
		return nil, err
	}

	ticketRes := &entity.TicketRes{
		ID:        ticket.ID,
		EventID:   ticket.EventID,
		UserID:    ticket.UserID,
		Status:    ticket.Status,
		CreatedAt: ticket.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: ticket.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ticketRes, nil
}

func (s *ticketService) FindAllTickets() ([]entity.TicketRes, error) {
	tickets, err := s.ticketRepository.FindAllTickets()
	if err != nil {
		return nil, err
	}

	var ticketRes []entity.TicketRes
	for _, ticket := range tickets {
		ticketRes = append(ticketRes, entity.TicketRes{
			ID:        ticket.ID,
			EventID:   ticket.EventID,
			UserID:    ticket.UserID,
			Status:    ticket.Status,
			CreatedAt: ticket.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: ticket.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return ticketRes, nil
}

func (s *ticketService) UpdateTicket(id int, req *entity.UpdateTicketReq) error {
	existingTicket, err := s.ticketRepository.FindTicketByID(id)
	if err != nil {
		return err
	}

	if req.Status != "" && req.Status != existingTicket.Status {
		existingTicket.Status = req.Status
	}

	return s.ticketRepository.UpdateTicket(id, existingTicket)
}

func (s *ticketService) DeleteTicket(id int) error {
	return s.ticketRepository.DeleteTicket(id)
}

func (s *ticketService) FindAllTicketsByUserID(userID int) ([]entity.TicketRes, error) {
	tickets, err := s.ticketRepository.FindAllTicketsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var ticketRes []entity.TicketRes
	for _, ticket := range tickets {
		ticketRes = append(ticketRes, entity.TicketRes{
			ID:        ticket.ID,
			EventID:   ticket.EventID,
			UserID:    ticket.UserID,
			Status:    ticket.Status,
			CreatedAt: ticket.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: ticket.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return ticketRes, nil
}
