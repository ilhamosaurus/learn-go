package models

import "time"

type Dealership struct {
	DealerID  string    `json:"dealer_id" gorm:"primaryKey;not null;unique"`
	DCode     string    `json:"code" gorm:"not null;unique"`
	Name      string    `json:"name" gorm:"not null;type:varchar(100)"`
	Address   string    `json:"address" gorm:"not null;type:varchar(100)"`
	Cars      []Car     `json:"cars" gorm:"foreignKey:DCode"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
