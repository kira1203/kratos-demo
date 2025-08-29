package mapper

import (
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	model "github.com/go-kratos/kratos-layout/internal/bondProductApi/data/grom/models"
	"github.com/shopspring/decimal"
)

// decimalToString converts decimal.NullDecimal to string.
func decimalToString(d decimal.NullDecimal) string {
	if d.Valid {
		return d.Decimal.String()
	}
	return ""
}

// ToBondProductItem converts model.TProduct to domain.BondProductItem.
func ToBondProductItem(do *model.TProduct) *domain.BondProductItem {
	if do == nil {
		return nil
	}
	return &domain.BondProductItem{
		ID:         do.Id,
		ProductId:  do.ProductId,
		Name:       do.Name,
		NameEn:     do.NameEn,
		Status:     do.Status,
		Isin:       do.Isin,
		IssuerType: do.IssuerType,
		Issuer:     do.Issuer,
		Currency:   do.Currency,
		MidYield:   decimalToString(do.MidYield),
	}
}
