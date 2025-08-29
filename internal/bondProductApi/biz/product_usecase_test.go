package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/dto"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockProductRepo 模拟 ProductRepo 接口
type MockProductRepo struct {
	OnBannerList func(context.Context) ([]*domain.BondProductItem, error)
}

func (m *MockProductRepo) BannerList(ctx context.Context) ([]*domain.BondProductItem, error) {
	if m.OnBannerList != nil {
		return m.OnBannerList(ctx)
	}
	return nil, nil
}

func (m *MockProductRepo) ProductList(ctx context.Context, req *dto.ProductQueryReq) ([]*domain.BondProductItem, error) {
	// 如果需要测试其他方法，也可以实现
	return nil, nil
}

// MockRepoRegistry 模拟 RepoRegistry 接口
type MockRepoRegistry struct {
	productFunc func() ProductRepo
}

func (m *MockRepoRegistry) Product() ProductRepo {
	if m.productFunc != nil {
		return m.productFunc()
	}
	return nil
}

func TestProductBiz_GetBannerList(t *testing.T) {
	// 给定：准备测试数据
	ctx := context.Background()
	expectedItems := []*domain.BondProductItem{
		{ID: 1, Name: "Bond A"},
		{ID: 2, Name: "Bond B"},
	}

	// 创建 mock
	mockProductRepo := &MockProductRepo{
		OnBannerList: func(ctx context.Context) ([]*domain.BondProductItem, error) {
			return expectedItems, nil
		},
	}

	mockRepoRegistry := &MockRepoRegistry{
		productFunc: func() ProductRepo {
			return mockProductRepo
		},
	}

	uc := NewProductBiz(nil, mockRepoRegistry, log.DefaultLogger)

	result, err := uc.GetBannerList(ctx)

	assert.NoError(t, err)           // 断言没有错误
	assert.NotNil(t, result)         // 断言 result 不为 nil
	assert.Equal(t, 0, result.Start) // 断言分页字段
	assert.Equal(t, 0, result.Count)
	assert.Equal(t, int64(0), result.Total)
	assert.Len(t, result.List, 2)               // 断言列表长度为 2
	assert.Equal(t, expectedItems, result.List) // 断言列表内容一致
}

func TestProductBiz_GetBannerList_Error(t *testing.T) {
	// 给定
	ctx := context.Background()
	mockError := errors.New("database error")

	mockProductRepo := &MockProductRepo{
		OnBannerList: func(ctx context.Context) ([]*domain.BondProductItem, error) {
			return nil, mockError
		},
	}

	mockRepoRegistry := &MockRepoRegistry{
		productFunc: func() ProductRepo {
			return mockProductRepo
		},
	}

	uc := NewProductBiz(nil, mockRepoRegistry, log.DefaultLogger)

	result, err := uc.GetBannerList(ctx)

	assert.Error(t, err)                        // 断言有错误
	assert.EqualError(t, err, "database error") // 断言错误信息
	assert.Nil(t, result)                       // 断言 result 为 nil
}
