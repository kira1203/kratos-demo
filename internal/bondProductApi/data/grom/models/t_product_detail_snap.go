package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type TProductDetailSnap struct {
	Id                            int64               `gorm:"column:id;primary_key;auto_increment;comment:'id'" json:"id"`
	ParentId                      int64               `gorm:"column:parent_id;comment:'t_product.id'" json:"parent_id"`
	RiskRating                    int                 `gorm:"column:risk_rating;comment:'风险评级.1-低风险,2-中低风险,3-中风险,4-中高风险,5-高风险'" json:"risk_rating"`
	InvestmentConcentration       int                 `gorm:"column:investment_concentration;comment:'投资集中度.1-<15%,2-<30%,3-<45%,4-<60%,5->=60%'" json:"investment_concentration"`
	TradeDirection                int                 `gorm:"column:trade_direction;comment:'交易方向.1-可买可卖,2-仅支持卖出'" json:"trade_direction"`
	MinBuyAmount                  decimal.NullDecimal `gorm:"column:min_buy_amount;type:decimal(19,4);comment:'最低买入金额'" json:"min_buy_amount"`
	IncrBuyAmount                 decimal.NullDecimal `gorm:"column:incr_buy_amount;type:decimal(19,4);comment:'递增买入金额'" json:"incr_buy_amount"`
	MinSellAmount                 decimal.NullDecimal `gorm:"column:min_sell_amount;type:decimal(19,4);comment:'最低卖出金额'" json:"min_sell_amount"`
	IncrSellAmount                decimal.NullDecimal `gorm:"column:incr_sell_amount;type:decimal(19,4);comment:'递增卖出金额'" json:"incr_sell_amount"`
	HasTradingFee                 int                 `gorm:"column:has_trading_fee;comment:'是否收取交易费.1-是,2-否'" json:"has_trading_fee"`
	TradingFee                    decimal.NullDecimal `gorm:"column:trading_fee;type:decimal(19,4);comment:'交易费'" json:"trading_fee"`
	AccruedInterestRate           decimal.NullDecimal `gorm:"column:accrued_interest_rate;type:decimal(19,4);comment:'代收取的利息费'" json:"accrued_interest_rate"`
	MinAccruedInterestFee         decimal.NullDecimal `gorm:"column:min_accrued_interest_fee;type:decimal(19,2);comment:'最低代收利息费'" json:"min_accrued_interest_fee"`
	MaxAccruedInterestFee         decimal.NullDecimal `gorm:"column:max_accrued_interest_fee;type:decimal(19,2);comment:'最高代收利息费'" json:"max_accrued_interest_fee"`
	OtherAccruedInterestFee       int                 `gorm:"column:other_accrued_interest_fee;comment:'其他代收利息费1是2否'" json:"other_accrued_interest_fee"`
	OtherAccruedInterestFeeRemark string              `gorm:"column:other_accrued_interest_fee_remark;type:varchar(255);comment:'其他代收利息费说明'" json:"other_accrued_interest_fee_remark"`
	HasCustodianFee               int                 `gorm:"column:has_custodian_fee;comment:'是否收取托管费.1-是,2-否'" json:"has_custodian_fee"`
	CustodianFee                  decimal.NullDecimal `gorm:"column:custodian_fee;type:decimal(19,4);comment:'托管费'" json:"custodian_fee"`
	CreateTime                    *time.Time          `gorm:"column:create_time;comment:'创建时间;autoCreateTime'" json:"create_time"`
	UpdateTime                    *time.Time          `gorm:"column:update_time;comment:'更新时间;autoCreateTime'" json:"update_time"`
}

func (t *TProductDetailSnap) TableName() string {
	return "t_product_detail_snap"
}

const (
	FieldTProductDetailSnapId                            = "id"
	FieldTProductDetailSnapParentId                      = "parent_id"
	FieldTProductDetailSnapRiskRating                    = "risk_rating"
	FieldTProductDetailSnapInvestmentConcentration       = "investment_concentration"
	FieldTProductDetailSnapTradeDirection                = "trade_direction"
	FieldTProductDetailSnapMinBuyAmount                  = "min_buy_amount"
	FieldTProductDetailSnapIncrBuyAmount                 = "incr_buy_amount"
	FieldTProductDetailSnapMinSellAmount                 = "min_sell_amount"
	FieldTProductDetailSnapIncrSellAmount                = "incr_sell_amount"
	FieldTProductDetailSnapHasTradingFee                 = "has_trading_fee"
	FieldTProductDetailSnapTradingFee                    = "trading_fee"
	FieldTProductDetailSnapAccruedInterestRate           = "accrued_interest_rate"
	FieldTProductDetailSnapMinAccruedInterestFee         = "min_accrued_interest_fee"
	FieldTProductDetailSnapMaxAccruedInterestFee         = "max_accrued_interest_fee"
	FieldTProductDetailSnapOtherAccruedInterestFee       = "other_accrued_interest_fee"
	FieldTProductDetailSnapOtherAccruedInterestFeeRemark = "other_accrued_interest_fee_remark"
	FieldTProductDetailSnapHasCustodianFee               = "has_custodian_fee"
	FieldTProductDetailSnapCustodianFee                  = "custodian_fee"
	FieldTProductDetailSnapCreateTime                    = "create_time"
	FieldTProductDetailSnapUpdateTime                    = "update_time"
)
