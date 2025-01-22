package controller

import (
	"net/http"
	"strconv"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/response"
	"github.com/Ayyasy123/dibimbing-take-home-test/service"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) *EventController {
	return &EventController{eventService: eventService}
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	var req entity.CreateEventReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	eventRes, err := c.eventService.CreateEvent(&req)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create event", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusCreated, "Event created successfully", eventRes)
}

func (c *EventController) FindEventByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid event ID", err)
		return
	}

	eventRes, err := c.eventService.FindEventByID(id)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve event", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "Event retrieved successfully", eventRes)
}

func (c *EventController) FindAllEvents(ctx *gin.Context) {
	eventsRes, err := c.eventService.FindAllEvents()
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve events", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "Events retrieved successfully", eventsRes)
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	var req entity.UpdateEventReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err := c.eventService.UpdateEvent(&req)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update event", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "Event updated successfully", nil)
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid event ID", err)
		return
	}

	err = c.eventService.DeleteEvent(id)
	if err != nil {
		response.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete event", err)
		return
	}

	response.SendSuccessResponse(ctx, http.StatusOK, "Event deleted successfully", nil)
}
