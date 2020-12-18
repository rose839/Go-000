package data

import (
	"week04/internal/do"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type OrderRepo struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewOrderRepo(db *gorm.DB, cache *redis.Client) *OrderRepo {
	return &OrderRepo{
		db:    db,
		cache: cache,
	}
}

func (d *OrderRepo) SaveOrder(order *do.Order) error {
	if err := d.db.Model(order).Create(order).Error; err != nil {
		return errors.Wrapf(err, "data save order failed! error: %s\n", err.Error())
	}
	return nil
}
