package model

import (
	"time"
)

type TProductFile struct {
	Id         int64      `gorm:"column:id;primary_key;auto_increment;comment:id" json:"id"`
	PID        int64      `gorm:"column:p_id;comment:产品id" json:"p_id"`
	Name       string     `gorm:"column:name;comment:文件名" json:"name"`
	OrigName   string     `gorm:"column:orig_name;comment:原始文件名" json:"orig_name"`
	FileId     string     `gorm:"column:file_id;comment:文件id" json:"file_id"`
	Suffix     string     `gorm:"column:suffix;type:varchar(32);comment:文件后缀" json:"suffix"`
	Status     int        `gorm:"column:status;type:tinyint(4);comment:状态.1-有效,2-无效" json:"status"`
	CreateBy   string     `gorm:"column:create_by;type:varchar(32);comment:创建人" json:"create_by"`
	Source     int        `gorm:"column:source;comment:文件来源1-星路 2-陆浦" json:"source"`
	CreateTime *time.Time `gorm:"column:create_time;comment:创建时间;autoCreateTime" json:"create_time"`
	UpdateBy   string     `gorm:"column:update_by;type:varchar(32);comment:更新人" json:"update_by"`
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间;autoCreateTime;autoUpdateTime" json:"update_time"`
}

func (t *TProductFile) TableName() string {
	return "t_product_file"
}

const (
	FieldTProductFileStatus = "status"

	SourceLUPU = 2
)

func (t *TProductFile) GetReadOnly() bool {
	if t.Source == SourceLUPU {
		return false
	}

	return true
}
