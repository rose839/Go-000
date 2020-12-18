package service

import (
	"week04/internal/biz"
	"context"
	pb "week04/api"
	"week04/internal/do"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Service struct {
	orderUserCase *biz.OrderCase
	pb.UnimplementedSaveOrderServer
}

// NewService is used to create a new service.
func NewService(orderUserCase *biz.OrderCase) *Service {
	return &Service{orderUserCase: orderUserCase}
}

func (s *service) SaveOrder(ctx context.Context, in *pb.Order, opts ...grpc.CallOption) (*pb.Resp, error) {
	o := &do.Order{}
	o.ID = in.GetId()
	o.Name = in.GetName()

	if err := o.Check(); err != nil {
		return nil, errors.WithMessage(err, "save order failed")
	}

	if s.orderUserCase.
}
