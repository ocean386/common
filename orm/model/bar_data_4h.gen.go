// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBarData4h = "bar_data_4h"

// BarData4h 时间粒度表-每4小时
type BarData4h struct {
	ID                int64   `gorm:"column:id;type:bigint(20);primaryKey;comment:主键ID" json:"id"`                                                                // 主键ID
	CurrencyID        string  `gorm:"column:currency_id;type:varchar(20);not null;comment:币种ID" json:"currency_id"`                                               // 币种ID
	BeginTime         int64   `gorm:"column:begin_time;type:bigint(20);not null;comment:开始时间" json:"begin_time"`                                                  // 开始时间
	OpenPrice         float64 `gorm:"column:open_price;type:decimal(10,3);not null;default:0.000;comment:开盘价格" json:"open_price"`                                 // 开盘价格
	ClosePrice        float64 `gorm:"column:close_price;type:decimal(10,3);not null;default:0.000;comment:收盘价格" json:"close_price"`                               // 收盘价格
	HighPrice         float64 `gorm:"column:high_price;type:decimal(10,3);not null;default:0.000;comment:最高价格" json:"high_price"`                                 // 最高价格
	LowPrice          float64 `gorm:"column:low_price;type:decimal(10,3);not null;default:0.000;comment:最低价格" json:"low_price"`                                   // 最低价格
	BarType           string  `gorm:"column:bar_type;type:varchar(10);not null;default:4H;comment:时间粒度类型:默认值1m,1m/5m/15m/30m/1H/4H/12H/1D/1W/1M" json:"bar_type"` // 时间粒度类型:默认值1m,1m/5m/15m/30m/1H/4H/12H/1D/1W/1M
	SigType           int64   `gorm:"column:sig_type;type:tinyint(8);not null;comment:行情信号: 0-/, 1-上涨, 2-快速上涨, 3-爆涨, 4-下跌, 5-快速下跌, 6-爆跌" json:"sig_type"`         // 行情信号: 0-/, 1-上涨, 2-快速上涨, 3-爆涨, 4-下跌, 5-快速下跌, 6-爆跌
	PriceChange       float64 `gorm:"column:price_change;type:decimal(10,3);not null;default:0.000;comment:涨跌幅度" json:"price_change"`                             // 涨跌幅度
	PriceRange        float64 `gorm:"column:price_range;type:decimal(10,3);not null;default:0.000;comment:涨跌振幅" json:"price_range"`                               // 涨跌振幅
	PriceRangePrecent float64 `gorm:"column:price_range_precent;type:decimal(10,3);not null;default:0.000;comment:涨跌振幅比例" json:"price_range_precent"`             // 涨跌振幅比例
	TriggerTime       int64   `gorm:"column:trigger_time;type:bigint(20);not null;comment:触发时间" json:"trigger_time"`                                              // 触发时间
}

// TableName BarData4h's table name
func (*BarData4h) TableName() string {
	return TableNameBarData4h
}
