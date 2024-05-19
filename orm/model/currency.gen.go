// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCurrency = "currency"

// Currency 虚拟币币种表
type Currency struct {
	ID           int64   `gorm:"column:id;type:bigint(20);primaryKey;comment:主键ID" json:"id"`                                           // 主键ID
	CurrencyID   string  `gorm:"column:currency_id;type:varchar(10);not null;comment:币种ID" json:"currency_id"`                          // 币种ID
	CurrencyName string  `gorm:"column:currency_name;type:varchar(10);not null;comment:币种-名称" json:"currency_name"`                     // 币种-名称
	MinUnitPrice float64 `gorm:"column:min_unit_price;type:decimal(8,4);not null;default:0.0000;comment:最小单元-价格" json:"min_unit_price"` // 最小单元-价格
	CreatedAt    int64   `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                                      // 创建时间
	UpdatedAt    int64   `gorm:"column:updated_at;autoUpdateTime:mill;comment:更新时间" json:"updated_at"`                                  // 更新时间
	Sort         int64   `gorm:"column:sort;type:tinyint(8);not null;default:1;comment:序号" json:"sort"`                                 // 序号
}

// TableName Currency's table name
func (*Currency) TableName() string {
	return TableNameCurrency
}
