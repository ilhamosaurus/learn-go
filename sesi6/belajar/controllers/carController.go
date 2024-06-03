package controllers

import (
	"fmt"
	"net/http"

	"belajar/models"
	"belajar/services"

	"github.com/gin-gonic/gin"
)

func CreateCar(ctx *gin.Context) {
	var newCar models.Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dealership, err := services.GetDealershipByCode(newCar.DealerCode)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if dealership.DealerID == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Dealer code not found",
			"message":      fmt.Sprintf("dealer with code %v not found", newCar.DealerCode),
		})
		return
	}

	cars, err := services.GetCars()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	Car := models.Car{
		CarID:      fmt.Sprintf("C-%d", len(cars)+1),
		Brand:      newCar.Brand,
		Model:      newCar.Model,
		Price:      newCar.Price,
		DealerCode: newCar.DealerCode,
	}

	createdCars, err := services.CreateCars(Car)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Car created successfully",
		"data":    createdCars,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	var car models.Car

	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedCar, err := services.UpdateCar(carID, car.Brand, car.Model, car.Price, car.DealerCode)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if updatedCar.CarID == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v updated successfully", carID),
		"data":    updatedCar,
	})
}

func GetCarByID(ctx *gin.Context) {
	carID := ctx.Param("carID")

	carData, err := services.GetCarByID(carID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if carData.CarID == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carId := ctx.Param("carID")

	car, err := services.GetCarByID(carId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if car.CarID == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      fmt.Sprintf("Car with id %v not found", carId),
		})
		return
	}

	err = services.DeleteCar(carId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v deleted successfully", carId),
	})
}

func GetCars(ctx *gin.Context) {
	carDatas, err := services.GetCars()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if len(carDatas) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      "Car data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": carDatas,
	})
}
