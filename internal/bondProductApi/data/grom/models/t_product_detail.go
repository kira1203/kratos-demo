package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type TProductDetail struct {
	Id                int64      `gorm:"column:id;primary_key;auto_increment;comment:id" json:"id"`
	ParentId          int64      `gorm:"column:parent_id;comment:t_product.id" json:"parent_id"`
	Name              string     `gorm:"column:name;type:varchar(255);comment:产品名称" json:"name"`
	NameEn            string     `gorm:"column:name_en;type:varchar(255);comment:产品英文名称" json:"name_en"`
	IssuerType        int        `gorm:"column:issuer_type;comment:发行人类型.1-主权国家,2-政府,3-企业" json:"issuer_type"`
	Issuer            string     `gorm:"column:issuer;comment:发行人" json:"issuer"`
	BondType          string     `gorm:"column:bond_type;type:varchar(255);comment:债券类型" json:"bond_type"`
	IssueDate         *time.Time `gorm:"column:issue_date;type:date;comment:发行日" json:"issue_date"`
	MaturityDate      *time.Time `gorm:"column:maturity_date;comment:到期日" json:"maturity_date"`
	Perpetual         int        `gorm:"column:perpetual;comment:是否永续.1-是,2-否." json:"perpetual"`
	Currency          string     `gorm:"column:currency;type:varchar(8);comment:币种.USD,HKD,CNH" json:"currency"`
	IssuanceAmount    *string    `gorm:"column:issuance_amount;type:decimal(19,2);comment:发行总额度" json:"issuance_amount"`
	CouponType        int        `gorm:"column:coupon_type;type:int;comment:票息类型.1-固定,2-浮动,3-零息" json:"coupon_type"`
	CouponRate        *string    `gorm:"column:coupon_rate;type:decimal(19,4);comment:票面息率" json:"coupon_rate"`
	IsDividend        int        `gorm:"column:is_dividend;type:int;comment:是否派息.1-是,2-否" json:"is_distribute"`
	DividendFreq      int        `gorm:"column:dividend_freq;type:int;comment:派息频率.1-月度,2-季度,3-半年度,4-年度" json:"dividend_freq"`
	Priority          string     `gorm:"column:priority;type:varchar(255);comment:优先权" json:"priority"`
	PriorityType      int        `gorm:"column:priority_type;comment:'优先权 1-第一留置权 2-优先抵押 3-优先无抵押 4-优先次级 5-次级 6-低顺位次级 7-其他'" json:"priority_type"`
	ListOnExchange    int        `gorm:"column:list_on_exchange;comment:是否上市交易所.1-是,2-否" json:"list_on_exchange"`
	Exchange          string     `gorm:"column:exchange;type:varchar(125);comment:上市交易所" json:"exchange"`
	BondSPRating      string     `gorm:"column:bond_sp_rating;type:varchar(255);comment:标普债券评级" json:"bond_sp_rating"`
	BondMoodyRating   string     `gorm:"column:bond_moody_rating;type:varchar(255);comment:穆迪债券评级" json:"bond_moody_rating"`
	BondFitchRating   string     `gorm:"column:bond_fitch_rating;type:varchar(255);comment:惠誉债券评级" json:"bond_fitch_rating"`
	IssuerSPRating    string     `gorm:"column:issuer_sp_rating;type:varchar(255);comment:标普发行人评级" json:"issuer_sp_rating"`
	IssuerMoodyRating string     `gorm:"column:issuer_moody_rating;type:varchar(255);comment:穆迪发行人评级" json:"issuer_moody_rating"`
	IssuerFitchRating string     `gorm:"column:issuer_fitch_rating;type:varchar(255);comment:惠誉发行人评级" json:"issuer_fitch_rating"`

	// 适当性检查
	IsDerivative int `gorm:"column:is_derivative;type:int;comment:是否衍生品.1-是,2-否" json:"is_derivative"`
	RiskRating   int `gorm:"column:risk_rating;type:int;comment:风险评级.1-低风险,2-中低风险,3-中风险,4-中高风险,5-高风险" json:"risk_rating"`

	// 交易规则
	TradeDirection                 int     `gorm:"column:trade_direction;comment:交易方向.1-可买可卖,2-仅支持卖出" json:"trade_direction"`
	PP                             int     `gorm:"column:pp;comment:是否支持PP下单.1-是,2-否" json:"pp"`
	MinBuyAmount                   *string `gorm:"column:min_buy_amount;type:decimal(19,4);comment:最低买入金额" json:"min_buy_amount"`
	IncrBuyAmount                  *string `gorm:"column:incr_buy_amount;type:decimal(19,4);comment:递增买入金额" json:"incr_buy_amount"`
	MinSellAmount                  *string `gorm:"column:min_sell_amount;type:decimal(19,4);comment:最低卖出金额" json:"min_sell_amount"`
	IncrSellAmount                 *string `gorm:"column:incr_sell_amount;type:decimal(19,4);comment:递增卖出金额" json:"incr_sell_amount"`
	HasTradingFee                  int     `gorm:"column:has_trading_fee;comment:是否收取交易费.1-是,2-否" json:"has_trading_fee"`
	TradingFee                     *string `gorm:"column:trading_fee;type:decimal(19,4);comment:交易费" json:"trading_fee"`
	HasCustodianFee                int     `gorm:"column:has_custodian_fee;comment:是否收取托管费.1-是,2-否" json:"has_custodian_fee"`
	CustodianFee                   *string `gorm:"column:custodian_fee;type:decimal(19,4);comment:托管费" json:"custodian_fee"`
	BuySettlementDay1              int     `gorm:"column:buy_settlement_day1;comment:买入交收规则:买入交收时间(天)" json:"buy_settlement_day1"`
	BuySettlementDay2              int     `gorm:"column:buy_settlement_day2;comment:买入交收规则:卖出交收时间(天)" json:"buy_settlement_day2"`
	HasOtherBuySettlementRule      int     `gorm:"column:has_other_buy_settlement_rule;comment:买入交收规则:其他是否勾选.1-是,2-否" json:"has_other_buy_settlement_rule"`
	OtherBuySettlementRule         string  `gorm:"column:other_buy_settlement_rule;type:varchar(255);comment:买入交收规则:其他" json:"other_buy_settlement_rule"`
	SellSettlementDay              int     `gorm:"column:sell_settlement_day;comment:卖出交收规则:卖出交收时间(天)" json:"sell_settlement_day"`
	HasOtherSellSettlementRule     int     `gorm:"column:has_other_sell_settlement_rule;comment:卖出交收规则:其他是否勾选.1-是,2-否" json:"has_other_sell_settlement_rule"`
	OtherSellSettlementRule        string  `gorm:"column:other_sell_settlement_rule;type:varchar(255);comment:卖出交收规则:其他" json:"other_sell_settlement_rule"`
	DividendSettlementDay          int     `gorm:"column:dividend_settlement_day;comment:派息交收时间(天)" json:"dividend_settlement_day"`
	HasOtherDividendSettlementRule int     `gorm:"column:has_other_dividend_settlement_rule;comment:买入交收规则:其他是否勾选.1-是,2-否" json:"has_other_dividend_settlement_rule"`
	OtherDividendSettlementRule    string  `gorm:"column:other_dividend_settlement_rule;type:varchar(255);comment:买入交收规则:其他" json:"other_dividend_settlement_rule"`
	MaturitySettlementDay          int     `gorm:"column:maturity_settlement_day;comment:到期交收时间(天)" json:"maturity_settlement_day"`
	HasOtherMaturitySettlementRule int     `gorm:"column:has_other_maturity_settlement_rule;comment:到期交收规则:其他是否勾选.1-是,2-否" json:"has_other_maturity_settlement_rule"`
	OtherMaturitySettlementRule    string  `gorm:"column:other_maturity_settlement_rule;type:varchar(255);comment:到期交收规则:其他" json:"other_maturity_settlement_rule"`

	// 销售渠道
	ChannelRule       int        `gorm:"column:channel_rule;comment:销售渠道.1-全部可见,2-部分可见,3-部分不可见" json:"channel_rule"`
	ChannelVal        string     `gorm:"column:channel_val;type:varchar(255);comment:销售渠道值.逗号分割" json:"channel_val"`
	SupportCrossBoard int        `gorm:"column:support_cross_board;comment:是否支持跨境理财通.1-是,2-否" json:"support_cross_board"`
	CreateTime        *time.Time `gorm:"column:create_time;comment:创建时间;autoCreateTime" json:"create_time"`
	UpdateTime        *time.Time `gorm:"column:update_time;comment:更新时间;autoCreateTime" json:"update_time"`

	InterestAccrualDate           *time.Time          `gorm:"column:interest_accrual_date;comment:起息日 格式YYYY-MM-DD" json:"interest_accrual_date"`
	NextPaymentDate               *time.Time          `gorm:"column:next_payment_date;comment:下一派息日 格式YYYY-MM-DD" json:"next_payment_date"`
	CurrentAmountOutstanding      decimal.NullDecimal `gorm:"column:current_amount_outstanding;comment:剩余规模" json:"current_amount_outstanding"`
	Factor                        decimal.NullDecimal `gorm:"column:factor;comment:票面剩余比率 (0-1)" json:"factor"`
	PriceAtIssue                  decimal.NullDecimal `gorm:"column:price_at_issue;comment:发行价格" json:"price_at_issue"`
	CouponDeviationType           string              `gorm:"column:coupon_deviation_type;comment:不规则付息期" json:"coupon_deviation_type"`
	IsCallable                    int                 `gorm:"column:is_callable;comment:是否可回售 1是 2否 0未知" json:"is_callable"`
	IsPuttable                    int                 `gorm:"column:is_puttable;comment:是否可赎回 1是 2否 0未知" json:"is_puttable"`
	AccruedInterestRate           decimal.NullDecimal `gorm:"column:accrued_interest_rate;comment:代收利息费率" json:"accrued_interest_rate"`
	MinAccruedInterestFee         decimal.NullDecimal `gorm:"column:min_accrued_interest_fee;comment:最低代收利息费" json:"min_accrued_interest_fee"`
	MaxAccruedInterestFee         decimal.NullDecimal `gorm:"column:max_accrued_interest_fee;comment:最高代收利息费" json:"max_accrued_interest_fee"`
	OtherAccruedInterestFee       int                 `gorm:"column:other_accrued_interest_fee;comment:其他代收利息费 1是 2否" json:"other_accrued_interest_fee"`
	OtherAccruedInterestFeeRemark string              `gorm:"column:other_accrued_interest_fee_remark;comment:其他代收利息费说明" json:"other_accrued_interest_fee_remark"`
	DefaultDate                   *time.Time          `gorm:"column:default_date;comment:违约日期" json:"default_date"`
	DefaultType                   string              `gorm:"column:default_type;comment:违约类型" json:"default_type"`
	BondPortfolio                 int                 `gorm:"column:bond_portfolio;comment:债券组合 1普通债券 2PI债券" json:"bond_portfolio"`
	InvestmentConcentration       int                 `gorm:"column:investment_concentration;comment:投资集中度 1 :< 15%  2: < 30%  3 :< 45%   4: < 60%   5: > 60%" json:"investment_concentration"`
	InvestYear                    int                 `gorm:"column:invest_year;comment:投资年限 1-1年以下 2-1-5年 3-6-10年 4-11-20年 5-20年以上" json:"invest_year"`
	InvestObjective               int                 `gorm:"column:invest_objective;comment:投资目标 1-保守为主 2-收入主导 3-收入及增长 4-增长主导 5-积极增长" json:"invest_objective"`
}

func (detail *TProductDetail) TableName() string {
	return "t_product_detail"
}

func (detail *TProductDetail) GetBondRating() string {
	if detail.BondSPRating != "" {
		return detail.BondSPRating
	}
	if detail.BondMoodyRating != "" {
		return detail.BondMoodyRating
	}
	if detail.BondFitchRating != "" {
		return detail.BondFitchRating
	}

	return ""
}

func (detail *TProductDetail) GetPriority() string {
	switch detail.PriorityType {
	case 1:
		return "第一留置权"
	case 2:
		return "优先抵押"
	case 3:
		return "优先无抵押"
	case 4:
		return "优先次级"
	case 5:
		return "次级"
	case 6:
		return "低顺位次级"
	case 7:
		return detail.Priority
	}

	return ""
}

const (
	FieldTProductDetailId = "id"
)
