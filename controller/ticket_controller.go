package controller

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/helper"
	"github.com/Ayyasy123/dibimbing-take-home-test/service"
	"github.com/gin-gonic/gin"
)

type TicketController struct {
	ticketService service.TicketService
}

func NewTicketController(ticketService service.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var req entity.CreateTicketReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	ticketRes, err := c.ticketService.CreateTicket(&req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create ticket", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusCreated, "Ticket created successfully", ticketRes)
}

func (c *TicketController) FindTicketByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid ticket ID", err)
		return
	}

	ticketRes, err := c.ticketService.FindTicketByID(id)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve ticket", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Ticket retrieved successfully", ticketRes)
}

func (c *TicketController) FindAllTickets(ctx *gin.Context) {
	var paginationReq helper.PaginationRequest
	if err := ctx.ShouldBindQuery(&paginationReq); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid pagination parameters", err)
		return
	}

	ticketsRes, err := c.ticketService.FindAllTickets()
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve tickets", err)
		return
	}

	var data []interface{}
	for _, ticket := range ticketsRes {
		data = append(data, ticket)
	}

	paginatedResponse := helper.Paginate(data, paginationReq.Page, paginationReq.Limit)

	helper.SendSuccessResponse(ctx, http.StatusOK, "Tickets retrieved successfully", paginatedResponse)
}

func (c *TicketController) UpdateTicket(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid ticket ID", err)
		return
	}

	var req entity.UpdateTicketReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err = c.ticketService.UpdateTicket(id, &req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update ticket", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Ticket updated successfully", nil)
}

func (c *TicketController) DeleteTicket(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid ticket ID", err)
		return
	}

	err = c.ticketService.DeleteTicket(id)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete ticket", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Ticket deleted successfully", nil)
}

func (c *TicketController) FindAllTicketsByUserID(ctx *gin.Context) {
	// ambil user id dari token
	userId, exists := ctx.Get("user_id")
	if !exists {
		helper.SendErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	userIdInt, ok := userId.(int)
	if !ok {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to parse user ID", nil)
		return
	}

	ticketsRes, err := c.ticketService.FindAllTicketsByUserID(userIdInt)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve tickets", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Tickets retrieved successfully", ticketsRes)
}
