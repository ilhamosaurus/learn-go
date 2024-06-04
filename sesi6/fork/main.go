package main

import (
	"log"

	"fork/handler"
	"fork/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	engine := gin.New()

	connectionString := "host=localhost user=postgres password=123 dbname=backend_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return
	}

	userRepo := &repository.UserRepo{DB: db}
	orderRepo := &repository.OrderRepo{DB: db}

	userRepo.Migrate()
	orderRepo.Migrate()

	userHdl := &handler.UserHdl{Repository: userRepo}
	orderHdl := &handler.OrderHdl{Repository: orderRepo}

	userGroup := engine.Group("/users")
	{
		userGroup.GET("", userHdl.GetGorm)
		userGroup.POST("", userHdl.CreateGorm)
	}

	orderGroup := engine.Group("/orders")
	{
		orderGroup.GET("", orderHdl.GetGorm)
		orderGroup.POST("", orderHdl.CreateGorm)
	}

	err = engine.Run(":5004")
	if err != nil {
		log.Fatal("Error running server: ", err)
		return
	}
}
