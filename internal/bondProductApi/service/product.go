package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/go-kratos/kratos-layout/gen/api/product/v1"
)

// BondProductService is a greeter service.
type ProductService struct {
	v1.UnimplementedProductServer

	product *biz.ProductBiz
}

// NewGreeterService new a greeter service.
func NewBondProductService(product *biz.ProductBiz) *ProductService {
	return &ProductService{
		product: product,
	}
}

func (s *ProductService) HttpRegister(httpSrv *http.Server) {
	v1.RegisterProductHTTPServer(httpSrv, s)
}

func (s *ProductService) GrpcRegister(grpcSrv *grpc.Server) {
	v1.RegisterProductServer(grpcSrv, s)
}

// SayHello implements helloworld.GreeterServer.
func (s *ProductService) Banner(ctx context.Context, in *v1.BannerRequest) (*v1.BannerReply, error) {
	list, err := s.product.GetBannerList(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("==", list)

	return &v1.BannerReply{}, nil
}
