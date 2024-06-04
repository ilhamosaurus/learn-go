package model

import "errors"

type Order struct {
	ID         uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     uint64 `json:"user_id" gorm:"column:user_id"`
	OrderName  string `json:"order_name" gorm:"column:order_name"`
	UserDetail *User  `json:"user_detail" gorm:"foreignKey:UserID"`
}

func (o *Order) Validate() error {
	if o.OrderName == "" {
		return errors.New("order name cannot be empty")
	}

	if o.UserID == 0 {
		return errors.New("user id cannot be empty")
	}
	return nil
}
