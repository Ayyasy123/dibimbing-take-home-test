package repository

import (
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
