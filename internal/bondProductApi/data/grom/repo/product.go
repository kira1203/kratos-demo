package repo

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	gormtool "github.com/go-kratos/kratos-layout/pkg/gorm"
	"github.com/go-kratos/kratos/v2/log"
)

type productRepo struct {
	gormtool.Sess
	log *log.Helper
}

// NewProductRepo .
func NewGormProductRepo(tx gormtool.Sess, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		tx,
		log.NewHelper(logger),
	}
}

func (r *productRepo) BannerList(context.Context) ([]*domain.BondProductItem, error) {
	var (
		res = []*domain.BondProductItem{}
	)

	r.Tx().Model(domain.BondProductItem{}).Find(&res)
	return res, nil
}
