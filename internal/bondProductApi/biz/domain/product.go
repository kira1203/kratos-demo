package domain

type BondProductItem struct {
	Id              int64  `json:"id" `                //
	ProductId       string `json:"product_id" `        // product_id
	Symbol          string `json:"symbol" `            // symbol
	Isin            string `json:"isin" `              // ISIN
	Name            string `json:"name" `              // 债券名称
	NameEn          string `json:"name_en" `           // 债券英文名称
	IssuerType      int    `json:"issuer_type" `       // 发行人类型 1-主权国家,2-政府,3-企业
	Status          int    `json:"status" `            // 发布状态.1-草稿,2-已提交,3-已上架,4-已驳回,5-已下架,10-删除
	Currency        string `json:"currency" `          // 币种
	MidYield        string `json:"mid_yield" `         // 参考到期收益率, 取自债券行情—当日行情 最新的参考中间价对应的到期收益率
	MinBuyAmount    string `json:"min_buy_amount" `    // 起投金额
	BondRating      string `json:"bond_rating" `       // 债券评级
	BondSPRating    string `json:"bond_sp_rating" `    // 债券评级
	BondMoodyRating string `json:"bond_moody_rating" ` // 债券评级
	BondFitchRating string `json:"bond_fitch_rating" ` // 债券评级
	RemainYear      int64  `json:"remain_year" `       // 剩余年数
	RemainDays      int64  `json:"remain_days" `       // 剩余天数
	Issuer          string `json:"issuer" `            // 发行主体名称
}
