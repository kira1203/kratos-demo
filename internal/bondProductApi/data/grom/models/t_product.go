package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type TProduct struct {
	Id                int64               `gorm:"column:id;primary_key;auto_increment;comment:id" json:"id"`
	ProductId         string              `gorm:"column:product_id;comment:产品id lupu侧;type:varchar(128)" json:"product_id"`
	SrcId             string              `gorm:"column:src_id;comment:产品id 星路侧;type:varchar(128)" json:"src_id"`
	Name              string              `gorm:"column:name;type:varchar(255);comment:产品名称" json:"name"`
	NameEn            string              `gorm:"column:name_en;type:varchar(255);comment:产品英文名称" json:"name_en"`
	Status            int                 `gorm:"column:status;comment:状态.1-草稿,2-已提交,3-已上架,4-已驳回,5-已下架,10-删除" json:"status"`
	Isin              string              `gorm:"column:isin;type:varchar(16);comment:ISIN" json:"isin"`
	IssuerType        int                 `gorm:"column:issuer_type;comment:发行人类型.1-主权国家,2-政府,3-企业" json:"issuer_type"`
	Issuer            string              `gorm:"column:issuer;comment:发行人" json:"issuer"`
	Currency          string              `gorm:"column:currency;type:varchar(8);comment:币种" json:"currency"`
	IssuanceAmount    *string             `gorm:"column:issuance_amount;type:decimal(19,2);comment:发行总额度" json:"issuance_amount"`
	SupportCrossBoard int                 `gorm:"column:support_cross_board;comment:是否支持跨境理财通.1-是,2-否" json:"support_cross_board"`
	CouponRate        *string             `gorm:"column:coupon_rate;type:decimal(19,4);comment:票面息率" json:"coupon_rate"`
	TradeDirection    int                 `gorm:"column:trade_direction;comment:交易方向.1-可买可卖,2-仅支持卖出" json:"trade_direction"`
	PublishBy         string              `gorm:"column:publish_by;type:varchar(32);comment:上架人" json:"publish_by"`
	PublishTime       *time.Time          `gorm:"column:publish_time;comment:上架时间" json:"publish_time"`
	AuditBy           string              `gorm:"column:audit_by;type:varchar(32);comment:审核人" json:"audit_by"`
	AuditTime         *time.Time          `gorm:"column:audit_time;comment:审核时间" json:"audit_time"`
	Top               int                 `gorm:"column:top;type:int;comment:是否置顶.1-是,2-否;default:2" json:"top"`
	TopTime           *time.Time          `gorm:"column:top_time;comment:置顶时间" json:"top_time"`
	IsDefault         int                 `gorm:"column:is_default;comment:是否违约" json:"is_default"`
	CreateBy          string              `gorm:"column:create_by;type:varchar(32);comment:创建人" json:"create_by"`
	CreateTime        *time.Time          `gorm:"column:create_time;comment:创建时间;autoCreateTime" json:"create_time"`
	UpdateBy          string              `gorm:"column:update_by;type:varchar(32);comment:更新人" json:"update_by"`
	UpdateTime        *time.Time          `gorm:"column:update_time;comment:更新时间;autoCreateTime" json:"update_time"`
	MidYield          decimal.NullDecimal `gorm:"column:mid_yield;comment:中间价对应的收益率" json:"mid_yield"`
	BidYield          decimal.NullDecimal `gorm:"column:bid_yield;comment:买盘价对应的收益率" json:"bid_yield"`
	AskYield          decimal.NullDecimal `gorm:"column:ask_yield;comment:卖盘价对应的收益率" json:"ask_yield"`
	ReferenceMidYield decimal.NullDecimal `gorm:"column:reference_mid_yield;comment:历史行情最新中间价收益率" json:"reference_mid_yield"`
	Files             []*TProductFile     `gorm:"foreignKey:PID;References:Id" json:"files"`
	Detail            *TProductDetail     `gorm:"foreignKey:ParentId;References:Id" json:"detail"`
	DetailSnap        *TProductDetailSnap `gorm:"foreignKey:ParentId;References:Id" json:"detail_snap"`
}

func (t *TProduct) TableName() string {
	return "t_product"
}

const (
	FieldTProductId         = "t_product.id"
	FieldTProductProductId  = "t_product.product_id"
	FieldTProductSrcId      = "t_product.src_id"
	FieldTProductName       = "t_product.name"
	FieldTProductNameEn     = "t_product.name_en"
	FieldTProductIssuerType = "t_product.issuer_type"
	FieldTProductIssuer     = "t_product.issuer"
	FieldTProductStatus     = "t_product.status"
	FieldTProductIsin       = "t_product.isin"

	//导航属性
	FieldTProductFiles         = "Files"
	FieldTProductDetail        = "Detail"
	FieldTProductDetailSnap    = "DetailSnap"
	FieldTProductPriceRealTime = "PriceRealTimeDetail"
)

type ProductRec struct {
	TProduct
}
