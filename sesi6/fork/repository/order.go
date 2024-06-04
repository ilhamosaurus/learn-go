package repository

import (
	"log"

	"fork/model"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func (o *OrderRepo) Migrate() {
	err := o.DB.AutoMigrate(&model.Order{})
	if err != nil {
		log.Fatal("Error migrating order: ", err)
	}
}

func (o *OrderRepo) Create(order *model.Order) error {
	err := o.DB.Debug().Model(&model.Order{}).Create(order).Error
	return err
}

func (o *OrderRepo) Get() ([]*model.Order, error) {
	orders := []*model.Order{}
	err := o.DB.Debug().Model(&model.Order{}).Preload("UserDetail").Find(&orders).Error
	return orders, err
}
