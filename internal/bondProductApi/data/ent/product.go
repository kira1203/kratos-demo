package ent

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	gormtool "github.com/go-kratos/kratos-layout/pkg/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type productRepo struct {
	sess gormtool.Sess
	log  *log.Helper
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
