package builder

import (
	"gorm.io/gorm"
	"kratos-demo/internal/bondProductApi/biz/dto"
	model "kratos-demo/internal/bondProductApi/data/grom/models"
)

// Build 返回 GORM Scope
func FromProductListReq(req *dto.ProductQueryReq) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if req.ID != 0 {
			db = db.Where("id = ?", req.ID)
		}

		if req.Paginate && req.Limit > 0 {
			db = db.Offset(req.Offset).Limit(req.Limit)
		}

		return db
	}
}

func FromUpdateReq(req *dto.ProductUpdateReq) map[string]interface{} {
	updates := make(map[string]interface{})

	if req.ID != 0 {
		updates[model.FieldTProductId] = req.ID
	}

	// 自动更新时间（可由 GORM Hook 处理）
	// updates["update_time"] = time.Now()

	return updates
}

/*
使用：
updater := FromUpdateReqWithExpr(req,
func(db *gorm.DB) *gorm.DB {
	return db.Update("view_count", gorm.Expr("view_count + 1"))
},
)

db.Scopes(updater)*/

func FromUpdateReqWithExpr(req *dto.ProductUpdateReq, expr ...func(db *gorm.DB) *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		updates := FromUpdateReq(req)
		if len(updates) > 0 {
			db = db.Updates(updates)
		}
		for _, f := range expr {
			db = f(db)
		}
		return db
	}
}
