package biz

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/dto"
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

func (uc *ProductBiz) GetBannerList(ctx context.Context) (*dto.ListRespData[domain.BondProductItem], error) {
	uc.log.WithContext(ctx).Infof("GetBannerList: %v", "===")
	list, err := uc.repo.Product().BannerList(ctx)
	if err != nil {
		return nil, err
	}
	res := &dto.ListRespData[domain.BondProductItem]{
		Start: 0,
		Count: 0,
		Total: 0,
		List:  list,
	}

	return res, nil
}
