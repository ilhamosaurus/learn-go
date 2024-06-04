package handler

import (
	"net/http"

	"fork/model"
	"fork/repository"

	"github.com/gin-gonic/gin"
)

type UserHdl struct {
	Repository *repository.UserRepo
}

func (u *UserHdl) GetGorm(ctx *gin.Context) {
	users, err := u.Repository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	if len(users) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (u *UserHdl) CreateGorm(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.BindJSON(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}

	if err := user.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	err := u.Repository.Create(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}
