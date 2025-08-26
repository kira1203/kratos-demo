package service

import (
	"context"
	v1 "github.com/go-kratos/kratos-layout/gen/api/product/v1"
	"github.com/go-kratos/kratos-layout/internal/user/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedProductServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) Banner(ctx context.Context, in *v1.BannerRequest) (*v1.BannerReply, error) {
	/*g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}*/
	return &v1.BannerReply{}, nil
}
