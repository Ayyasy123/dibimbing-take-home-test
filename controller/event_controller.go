package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"github.com/Ayyasy123/dibimbing-take-home-test/helper"
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
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Perform validation
	if err := helper.ValidateStruct(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	eventRes, err := c.eventService.CreateEvent(&req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to create event", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusCreated, "Event created successfully", eventRes)
}

func (c *EventController) FindEventByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid event ID", err)
		return
	}

	eventRes, err := c.eventService.FindEventByID(id)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve event", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Event retrieved successfully", eventRes)
}

func (c *EventController) FindAllEvents(ctx *gin.Context) {
	var paginationReq helper.PaginationRequest
	if err := ctx.ShouldBindQuery(&paginationReq); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid pagination parameters", err)
		return
	}

	// Set default values if page or limit is not provided
	if paginationReq.Page == 0 {
		paginationReq.Page = 1 // Default page is 1
	}
	if paginationReq.Limit == 0 {
		paginationReq.Limit = 10 // Default limit is 10
	}

	eventsRes, err := c.eventService.FindAllEvents()
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve events", err)
		return
	}

	var data []interface{}
	for _, event := range eventsRes {
		data = append(data, event)
	}

	paginatedResponse := helper.Paginate(data, paginationReq.Page, paginationReq.Limit)

	helper.SendSuccessResponse(ctx, http.StatusOK, "Events retrieved successfully", paginatedResponse)
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid event ID", err)
		return
	}

	var req entity.UpdateEventReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Perform validation
	if err := helper.ValidateStruct(&req); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	err = c.eventService.UpdateEvent(id, &req)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to update event", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Event updated successfully", nil)
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid event ID", err)
		return
	}

	err = c.eventService.DeleteEvent(id)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete event", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Event deleted successfully", nil)
}

func (c *EventController) SearchEvents(ctx *gin.Context) {
	// TODO: Implement event search logic
	searchQuery := ctx.Query("search")
	minPriceStr := ctx.Query("min_price")
	maxPriceStr := ctx.Query("max_price")
	category := ctx.Query("category")
	status := ctx.Query("status")
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	var minPrice, maxPrice int
	var startDate, endDate time.Time
	var err error

	if minPriceStr == "" {
		minPrice = 0
	} else {
		minPrice, err = strconv.Atoi(minPriceStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid min_price", err)
			return
		}
	}

	if maxPriceStr == "" {
		maxPrice = 100000000
	} else {
		maxPrice, err = strconv.Atoi(maxPriceStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid max_price", err)
			return
		}
	}

	if minPrice > maxPrice {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "min_price must be less than or equal to max_price", nil)
		return
	}

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid start_date", err)
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid end_date", err)
			return
		}
	}

	eventsRes, err := c.eventService.SearchEvents(searchQuery, minPrice, maxPrice, category, status, startDate, endDate)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve events", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Events retrieved successfully", eventsRes)
}

func (c *EventController) GetEventReport(ctx *gin.Context) {
	// Ambil start_date dan end_date dari query string (opsional)
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	var startDate, endDate time.Time
	var err error

	// Parse start_date jika ada
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid start_date format (YYYY-MM-DD)", err)
			return
		}
	}

	// Parse end_date jika ada
	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid end_date format (YYYY-MM-DD)", err)
			return
		}
	}

	// Pastikan start_date tidak lebih besar dari end_date jika keduanya ada
	if !startDate.IsZero() && !endDate.IsZero() && startDate.After(endDate) {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "start_date must be before or equal to end_date", nil)
		return
	}

	// Panggil service untuk menghasilkan laporan
	report, err := c.eventService.GetEventReport(startDate, endDate)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate event report", err)
		return
	}

	helper.SendSuccessResponse(ctx, http.StatusOK, "Event report generated successfully", report)
}

func (c *EventController) CancelEvent(ctx *gin.Context) {
	// Ambil ID event dari URL parameter
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid event ID", err)
		return
	}

	// Panggil service untuk membatalkan event
	if err := c.eventService.CancelEvent(eventID); err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to cancel event", err)
		return
	}

	// Kirim response sukses
	helper.SendSuccessResponse(ctx, http.StatusOK, "Event and associated tickets cancelled successfully", nil)
}
