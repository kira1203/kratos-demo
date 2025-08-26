package dto

type ProductQueryReq struct {
	ID       int
	Paginate bool
	Offset   int
	Limit    int
}
