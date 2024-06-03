package services

import (
	"fmt"

	"belajar/database"
	"belajar/models"
)

func CreateDealership(name string, address string) (models.Dealership, error) {
	db := database.GetDB()

	dealerships, err := GetDealerships()
	if err != nil {
		fmt.Println("Error getting dealerships:", err)
		return models.Dealership{}, err
	}

	Dealership := models.Dealership{
		DealerID: fmt.Sprintf("D-%d", len(dealerships)+1),
		DCode:    fmt.Sprintf("%s-%d", name, len(dealerships)+1),
		Name:     name,
		Address:  address,
	}

	errors := db.Create(&Dealership).Error
	if errors != nil {
		fmt.Println("Error creating dealership:", errors)
		return models.Dealership{}, errors
	}

	return Dealership, nil
}

func GetDealerships() ([]models.Dealership, error) {
	db := database.GetDB()

	var dealerships []models.Dealership
	err := db.Find(&dealerships).Error
	if err != nil {
		fmt.Println("Error getting dealerships:", err)
		return nil, err
	}

	return dealerships, nil
}

func GetDealershipByID(id string) (models.Dealership, error) {
	db := database.GetDB()

	var dealership models.Dealership

	err := db.Where("dealer_id = ?", id).First(&dealership).Error
	if err != nil {
		fmt.Println("Error getting dealership:", err)
		return models.Dealership{}, err
	}

	return dealership, nil
}

func UpdateDealership(dealerId string, name string, address string) (models.Dealership, error) {
	db := database.GetDB()

	dealership := models.Dealership{}

	err := db.Model(&dealership).Where("dealer_id = ?", dealerId).Updates(models.Dealership{Name: name, Address: address}).Error
	if err != nil {
		fmt.Println("Error updating dealership:", err)
		return models.Dealership{}, err
	}

	return dealership, nil
}

func DeleteDealership(dealerId string) error {
	db := database.GetDB()

	err := db.Where("dealer_id = ?", dealerId).Delete(&models.Dealership{}).Error
	if err != nil {
		fmt.Println("Error deleting dealership:", err)
		return err
	}

	return nil
}

func GetDealershipByCode(d_code string) (models.Dealership, error) {
	db := database.GetDB()

	var dealership models.Dealership
	err := db.Where("d_code = ?", d_code).First(&dealership).Error
	if err != nil {
		fmt.Println("Error getting dealership by code: ", err)
		return models.Dealership{}, err
	}

	return dealership, nil
}
