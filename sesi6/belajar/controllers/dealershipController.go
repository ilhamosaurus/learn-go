package controllers

import (
	"net/http"

	"belajar/models"
	"belajar/services"

	"github.com/gin-gonic/gin"
)

func CreateDealership(ctx *gin.Context) {
	var newDealership models.Dealership

	if err := ctx.ShouldBindJSON(&newDealership); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	createdDealership, err := services.CreateDealership(newDealership.Name, newDealership.Address)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Dealership created successfully",
		"data":    createdDealership,
	})
}

func GetDealerships(ctx *gin.Context) {
	dealerships, err := services.GetDealerships()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(dealerships) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      "Dealership data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": dealerships,
	})
}

func GetDealershipByID(ctx *gin.Context) {
	dealershipID := ctx.Param("dealershipID")
	dealership, err := services.GetDealershipByID(dealershipID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if dealership.DealerID == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"message":      "Dealership data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": dealership,
	})
}

func UpdateDealership(ctx *gin.Context) {
	dealershipID := ctx.Param("dealershipID")
	var updatedDealership models.Dealership

	if err := ctx.ShouldBindJSON(&updatedDealership); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedDealership, err := services.UpdateDealership(dealershipID, updatedDealership.Name, updatedDealership.Address)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Dealership updated successfully",
		"data":    updatedDealership,
	})
}

func DeleteDealership(ctx *gin.Context) {
	dealershipID := ctx.Param("dealershipID")
	err := services.DeleteDealership(dealershipID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Dealership deleted successfully",
	})
}
