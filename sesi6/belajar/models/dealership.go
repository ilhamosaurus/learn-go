package models

import "time"

type Dealership struct {
	DealerID  string    `json:"dealer_id" gorm:"column:dealer_id;primaryKey;not null;unique"`
	DCode     string    `json:"d_code" gorm:"column:d_code;not null;unique"`
	Name      string    `json:"name" gorm:"column:name;not null;type:varchar(100)"`
	Address   string    `json:"address" gorm:"column:address;not null;type:varchar(100)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
