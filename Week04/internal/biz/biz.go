package biz

import (
	"week04/internal/data"
	"week04/internal/do"
)

type OrderRepoProc interface {
	SaveOrder(order *do.Order) error
}

type OrderCase struct {
	repo OrderRepoProc
}

func NewOrderBIZ(repo *data.OrderRepo) *OrderCase {
	return &OrderCase{repo: repo}
}

func (b *OrderCase) SaveOrder(order *do.Order) error {
	return b.repo.SaveOrder(order)
}
