package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods", controller.UpdateFood())

}
