package main

import (
	"week04/internal/biz"
	"week04/internal/data"
	"week04/internal/service"

	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitializeOrderService(db *gorm.DB, cache *redis.Client) *service.Service {
	wire.Build(service.NewService, biz.NewOrderBIZ, data.NewOrderRepo)
	return &service.Service{}
}
