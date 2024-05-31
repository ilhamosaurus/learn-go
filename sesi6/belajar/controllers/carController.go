package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Car created successfully",
		"data":    CarDatas,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}

	if !condition {
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
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
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
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carId == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      fmt.Sprintf("car with id %v not found", carId),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Car with id %v deleted successfully", carId),
	})
}

func GetCars(ctx *gin.Context) {
	if len(CarDatas) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      "Car data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": CarDatas,
	})
}
