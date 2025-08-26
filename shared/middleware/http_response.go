package middleware

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	// 实现统一返回逻辑
	return json.NewEncoder(w).Encode(APIResponse{
		Code:    0,
		Message: "OK",
		Result:  v,
	})
}
