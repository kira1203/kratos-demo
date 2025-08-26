package biz

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	"github.com/go-kratos/kratos/v2/log"
)

type ProductBiz struct {
	tx   TxManager
	repo RepoRegistry
	log  *log.Helper
}

func NewProductBiz(tx TxManager, repo RepoRegistry, logger log.Logger) *ProductBiz {
	return &ProductBiz{
		tx:   tx,
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ProductBiz) GetBannerList(ctx context.Context) ([]*domain.BondProductItem, error) {
	uc.log.WithContext(ctx).Infof("GetBannerList: %v", "===")
	return uc.repo.Product().BannerList(ctx)
}
