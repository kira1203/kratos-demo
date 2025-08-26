package dto

type ListRespData[T any] struct {
	Start int   `json:"start,omitempty"` // 分页起始位移
	Count int   `json:"count,omitempty"` // 分页大小
	Total int64 `json:"total,omitempty"` // 总数
	List  []*T  `json:"list"`            // 分页内容
}
