package data

import (
	"kratos-demo/internal/bondProductApi/biz"
)

type repoRegistry struct {
	productRepo biz.ProductRepo
}

func (r *repoRegistry) Product() biz.ProductRepo { return r.productRepo }

func NewRepoRegistry(product biz.ProductRepo) biz.RepoRegistry {
	return &repoRegistry{
		productRepo: product,
	}
}
