package handler

import (
	"fmt"
	"net/http"

	"fork/model"
	"fork/repository"

	"github.com/gin-gonic/gin"
)

type OrderHdl struct {
	Repository *repository.OrderRepo
}

func (o *OrderHdl) GetGorm(ctx *gin.Context) {
	orders, err := o.Repository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to get data: %s", err),
		})
		return
	}
	if len(orders) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (o *OrderHdl) CreateGorm(ctx *gin.Context) {
	order := &model.Order{}
	if err := ctx.BindJSON(order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body request",
		})
		return
	}

	if err := order.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	err := o.Repository.Create(order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}
