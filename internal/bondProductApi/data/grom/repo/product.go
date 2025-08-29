package repo

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/internal/bondProductApi/biz"
	"kratos-demo/internal/bondProductApi/biz/domain"
	"kratos-demo/internal/bondProductApi/biz/dto"
	"kratos-demo/internal/bondProductApi/data/grom/builder"
	gormtool "kratos-demo/pkg/gorm"
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

func (r *productRepo) ProductList(ctx context.Context, req *dto.ProductQueryReq) ([]*domain.BondProductItem, error) {
	var (
		res = []*domain.BondProductItem{}
	)

	r.Tx().Model(domain.BondProductItem{}).Scopes(builder.FromProductListReq(req)).Find(&res)

	return res, nil
}
