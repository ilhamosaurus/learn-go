package services

import (
	"fmt"
	"time"

	"belajar/database"
	"belajar/models"
)

func GetCars() ([]models.Car, error) {
	db := database.GetDB()

	var cars []models.Car
	err := db.Preload("Dealer").Find(&cars).Error
	if err != nil {
		fmt.Println("Error getting cars: ", err)
		return nil, err
	}

	return cars, nil
}

func CreateCars(car models.Car) (models.Car, error) {
	db := database.GetDB()

	errors := db.Create(&car).Error
	if errors != nil {
		fmt.Println("Failed to create Car: ", errors)
		return models.Car{}, errors
	}

	return car, nil
}

func GetCarByID(id string) (models.Car, error) {
	db := database.GetDB()

	var car models.Car

	err := db.Where("car_id = ?", id).Preload("Dealer").First(&car).Error
	if err != nil {
		fmt.Println("Error getting car: ", err)
		return models.Car{}, err
	}

	return car, nil
}

func UpdateCar(id string, brand string, model string, price int, d_code string) (models.Car, error) {
	db := database.GetDB()

	car, err := GetCarByID(id)
	if err != nil {
		return models.Car{}, err
	}

	dealership, err := GetDealershipByCode(d_code)
	if err != nil {
		return models.Car{}, err
	}

	time.Sleep(1 * time.Second)
	errors := db.Model(&car).Where("car_id = ?", id).Updates(models.Car{Brand: brand, Model: model, Price: price, DealerCode: dealership.DCode}).Error
	if errors != nil {
		fmt.Println("Failed to update car: ", errors)
		return models.Car{}, errors
	}

	return car, nil
}

func DeleteCar(id string) error {
	db := database.GetDB()

	err := db.Where("car_id = ?", id).Delete(&models.Car{}).Error
	if err != nil {
		fmt.Println("Error deleting car: ", err)
		return err
	}

	return nil
}
