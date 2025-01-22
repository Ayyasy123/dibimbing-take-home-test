package routes

import (
	"github.com/Ayyasy123/dibimbing-take-home-test/controller"
	"github.com/Ayyasy123/dibimbing-take-home-test/middleware"
	"github.com/Ayyasy123/dibimbing-take-home-test/repository"
	"github.com/Ayyasy123/dibimbing-take-home-test/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(db *gorm.DB, r *gin.Engine) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.LoginUser)

	userRoutes := r.Group("/users")
	userRoutes.Use(middleware.JWTAuth())
	{
		userRoutes.GET("", userController.FindAllUsers)
		userRoutes.GET("/:id", userController.FindUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}

func SetupEventRoutes(db *gorm.DB, r *gin.Engine) {
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventController := controller.NewEventController(eventService)

	eventRoutes := r.Group("/events")
	eventRoutes.Use(middleware.JWTAuth())
	{
		eventRoutes.POST("", eventController.CreateEvent)
		eventRoutes.GET("", eventController.FindAllEvents)
		eventRoutes.GET("/:id", eventController.FindEventByID)
		eventRoutes.PUT("/:id", eventController.UpdateEvent)
		eventRoutes.DELETE("/:id", eventController.DeleteEvent)
	}
}

func SetupTicketRoutes(db *gorm.DB, r *gin.Engine) {
	ticketRepo := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepo)
	ticketController := controller.NewTicketController(ticketService)

	ticketRoutes := r.Group("/tickets")
	ticketRoutes.Use(middleware.JWTAuth())
	{
		ticketRoutes.POST("", ticketController.CreateTicket)
		ticketRoutes.GET("", ticketController.FindAllTickets)
		ticketRoutes.GET("/:id", ticketController.FindTicketByID)
		ticketRoutes.PUT("/:id", ticketController.UpdateTicket)
		ticketRoutes.DELETE("/:id", ticketController.DeleteTicket)
	}
}
