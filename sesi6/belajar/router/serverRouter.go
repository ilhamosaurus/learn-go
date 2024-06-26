package router

import (
	"belajar/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controllers.GetCars)

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.GET("/cars/:carID", controllers.GetCarByID)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	router.GET("/dealers", controllers.GetDealerships)

	router.POST("/dealers", controllers.CreateDealership)

	router.PUT("/dealers/:dealershipID", controllers.UpdateDealership)

	router.GET("/dealers/:dealershipID", controllers.GetDealershipByID)

	router.DELETE("/dealers/:dealershipID", controllers.DeleteDealership)

	return router
}
