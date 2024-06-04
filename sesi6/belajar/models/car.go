package models

import "time"

type Car struct {
	CarID      string     `json:"car_id" gorm:"primaryKey"`
	Brand      string     `json:"brand" gorm:"not null;type:varchar(100)"`
	Model      string     `json:"model" gorm:"not null;type:varchar(100)"`
	Price      int        `json:"price" gorm:"not null"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
	DealerCode string     `json:"d_code"`
	Dealer     Dealership `json:"dealership" gorm:"foreignKey:DealerCode;references:DCode"`
}
