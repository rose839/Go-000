package do

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	ID   string `gorm:"column:name" json:"name" form:"name"`
	Name string `gorm:"column:password" json:"password" form:"password"`
}

func (o *Order) Check() error {
	if o.ID == "" {
		return errors.New("Order ID is empty")
	}
	if o.Name == "" {
		return errors.New("Order name is empty")
	}
	return nil
}
