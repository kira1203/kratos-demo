package ent

import (
	"context"
	"kratos-demo/internal/bondProductApi/biz"
	"kratos-demo/internal/bondProductApi/biz/domain"
	"kratos-demo/internal/bondProductApi/biz/dto"
	gormtool "kratos-demo/pkg/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type productRepo struct {
	sess gormtool.Sess
	log  *log.Helper
}

func (r *productRepo) ProductList(ctx context.Context, req *dto.ProductQueryReq) ([]*domain.BondProductItem, error) {
	//TODO implement me
	panic("implement me")
}

// NewProductRepo .
func NewEntProductRepo(tx gormtool.Sess, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		sess: tx,
		log:  log.NewHelper(logger),
	}
}

func (r *productRepo) BannerList(context.Context) ([]*domain.BondProductItem, error) {
	return nil, nil
}
