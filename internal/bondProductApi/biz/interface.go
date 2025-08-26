package biz

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/dto"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type RepoRegistry interface {
	Product() ProductRepo
}

// GreeterRepo is a Greater repo.
type ProductRepo interface {
	BannerList(context.Context) ([]*domain.BondProductItem, error)
	ProductList(context.Context, *dto.ProductQueryReq) ([]*domain.BondProductItem, error)
}

type TxManager interface {
	BeginTx(opts ...*sql.TxOptions) error
	RollbackTx()
	SavePoint(sp string) error
	CommitTx() error
	SetContext(ctx context.Context)
}
