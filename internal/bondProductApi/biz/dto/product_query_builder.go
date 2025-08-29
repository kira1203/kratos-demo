package dto

type ProductQueryReq struct {
	ID       int
	Paginate bool
	Offset   int
	Limit    int
}

type ProductUpdateReq struct {
	ID     int
	NameCN string
	NameTC string
	NameEN string
}
