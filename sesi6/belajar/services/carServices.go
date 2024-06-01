package services

import (
	"fmt"

	"belajar/database"
	"belajar/models"
)

func GetCars() ([]models.Car, error) {
	db := database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error
	if err != nil {
		fmt.Println("Error getting cars: ", err)
		return nil, err
	}

	return cars, nil
}

func CreateCars(brand string, model string, price int, d_code string) (models.Car, error) {
	db := database.GetDB()

	dealership, err := GetDealershipByCode(d_code)
	if err != nil {
		return models.Car{}, err
	}

	cars, err := GetCars()
	if err != nil {
		return models.Car{}, err
	}

	Car := models.Car{
		CarID: fmt.Sprintf("C-%d", len(cars)+1),
		Brand: brand,
		Model: model,
		Price: price,
		DCode: dealership.DCode,
	}

	errors := db.Create(&Car).Error
	if errors != nil {
		fmt.Println("Failed to create Car: ", errors)
		return models.Car{}, errors
	}

	return Car, nil
}
