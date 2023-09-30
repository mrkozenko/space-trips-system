package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrkozenko/space-trips-system/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		userTrips := api.Group("/user-trips")
		{
			//check continue add new route for get info by some types of places or spaceships
			userTrips.POST("/", h.addUserTrip)
			userTrips.GET("/", h.getAllUserTrips)
			userTrips.GET("/:id", h.getUserTripById)
			userTrips.PUT("/:id", h.updateUserTrip)
			userTrips.DELETE("/:id", h.deleteUserTrip)
		}

		spaceShips := api.Group("/spaceships")
		{
			spaceShips.POST("/", h.createSpaceship)
			spaceShips.GET("/", h.getAllSpaceships)
			spaceShips.GET("/:id", h.getSpaceTripById)
			spaceShips.PUT("/:id", h.updateSpaceship)
			spaceShips.DELETE("/:id", h.deleteSpaceship)
		}

		spaceTrips := api.Group("/space-trips")
		{
			//check continue add new route for get info by some types of places or spaceships
			spaceTrips.POST("/", h.createSpaceTrip)
			spaceTrips.GET("/", h.getAllSpaceTrips)
			spaceTrips.GET("/:id", h.getSpaceTripById)
			spaceTrips.PUT("/:id", h.updateSpaceTrip)
			spaceTrips.DELETE("/:id", h.deleteSpaceTrip)
		}
	}
	return router
}
