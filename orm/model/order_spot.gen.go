// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrderSpot = "order_spot"

// OrderSpot 现货订单列表
type OrderSpot struct {
	ID           int64   `gorm:"column:id;type:bigint(20);primaryKey;comment:主键ID" json:"id"`                                                          // 主键ID
	OrderPid     string  `gorm:"column:order_pid;type:varchar(20);not null;comment:OK平台订单ID" json:"order_pid"`                                         // OK平台订单ID
	OrderCid     string  `gorm:"column:order_cid;type:varchar(20);not null;comment:客户自定义订单ID" json:"order_cid"`                                        // 客户自定义订单ID
	UserID       int64   `gorm:"column:user_id;type:int(11);not null;default:10000;comment:用户ID" json:"user_id"`                                       // 用户ID
	CurrencyID   string  `gorm:"column:currency_id;type:varchar(10);not null;comment:币种ID" json:"currency_id"`                                         // 币种ID
	Simulation   int64   `gorm:"column:simulation;type:tinyint(8);not null;default:1;comment:模拟订单: 0-全部,1-否, 2-是" json:"simulation"`                   // 模拟订单: 0-全部,1-否, 2-是
	BuySellType  int64   `gorm:"column:buy_sell_type;type:tinyint(8);not null;comment:交易方向: 0-全部,1-买入,2-卖出" json:"buy_sell_type"`                      // 交易方向: 0-全部,1-买入,2-卖出
	OrderStatus  int64   `gorm:"column:order_status;type:tinyint(8);not null;comment:订单状态:0-全部,1-等待提交,2-等待成交,3-部分成交,4-完全成交,5-已撤销" json:"order_status"` // 订单状态:0-全部,1-等待提交,2-等待成交,3-部分成交,4-完全成交,5-已撤销
	OrderType    int64   `gorm:"column:order_type;type:tinyint(8);not null;comment:交易类型: 0-全部,1-高级限价,2-限价委托,3-市价委托 4-计划委托" json:"order_type"`          // 交易类型: 0-全部,1-高级限价,2-限价委托,3-市价委托 4-计划委托
	Total        float64 `gorm:"column:total;type:decimal(12,4);not null;default:0.0000;comment:委托总量" json:"total"`                                    // 委托总量
	OrderPrice   float64 `gorm:"column:order_price;type:decimal(12,4);not null;default:0.0000;comment:委托价格" json:"order_price"`                        // 委托价格
	TotalAmount  float64 `gorm:"column:total_amount;type:decimal(12,4);not null;default:0.0000;comment:委托金额" json:"total_amount"`                      // 委托金额
	Filled       float64 `gorm:"column:filled;type:decimal(12,4);not null;default:0.0000;comment:已成交量" json:"filled"`                                  // 已成交量
	FillPrice    float64 `gorm:"column:fill_price;type:decimal(12,4);not null;default:0.0000;comment:成交均价" json:"fill_price"`                          // 成交均价
	FilledValue  float64 `gorm:"column:filled_value;type:decimal(12,4);not null;default:0.0000;comment:已成交金额" json:"filled_value"`                     // 已成交金额
	ProfitAmount float64 `gorm:"column:profit_amount;type:decimal(12,4);not null;default:0.0000;comment:盈亏金额" json:"profit_amount"`                    // 盈亏金额
	Fee          float64 `gorm:"column:fee;type:decimal(12,6);not null;default:0.000000;comment:手续费" json:"fee"`                                       // 手续费
	CreatedAt    int64   `gorm:"column:created_at;comment:创建时间" json:"created_at"`                                                                     // 创建时间
	FilledAt     int64   `gorm:"column:filled_at;type:bigint(20);not null;comment:成交时间" json:"filled_at"`                                              // 成交时间
	CanceledAt   int64   `gorm:"column:canceled_at;type:bigint(20);not null;comment:撤销时间" json:"canceled_at"`                                          // 撤销时间
}

// TableName OrderSpot's table name
func (*OrderSpot) TableName() string {
	return TableNameOrderSpot
}
