// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStockConcept = "stock_concept"

// StockConcept 概念列表-A股
type StockConcept struct {
	ID             int64     `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`          // 主键ID
	ConceptCode    string    `gorm:"column:concept_code;type:varchar(10);not null;comment:概念代码" json:"concept_code"`                   // 概念代码
	ConceptName    string    `gorm:"column:concept_name;type:varchar(50);not null;comment:概念名称" json:"concept_name"`                   // 概念名称
	IsWatchConcept int64     `gorm:"column:is_watch_concept;type:tinyint(4);not null;comment:自选概念标志(0-否 1-是)" json:"is_watch_concept"` // 自选概念标志(0-否 1-是)
	IsDeleted      int64     `gorm:"column:is_deleted;type:tinyint(4);not null;comment:删除标记(0-否 1-是)" json:"is_deleted"`               // 删除标记(0-否 1-是)
	CreatedAt      time.Time `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                                 // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime:mill;comment:更新时间" json:"updated_at"`                             // 更新时间
}

// TableName StockConcept's table name
func (*StockConcept) TableName() string {
	return TableNameStockConcept
}
