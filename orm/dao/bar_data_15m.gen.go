// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ocean386/common/orm/model"
)

func newBarData15m(db *gorm.DB, opts ...gen.DOOption) barData15m {
	_barData15m := barData15m{}

	_barData15m.barData15mDo.UseDB(db, opts...)
	_barData15m.barData15mDo.UseModel(&model.BarData15m{})

	tableName := _barData15m.barData15mDo.TableName()
	_barData15m.ALL = field.NewAsterisk(tableName)
	_barData15m.ID = field.NewInt64(tableName, "id")
	_barData15m.CurrencyID = field.NewString(tableName, "currency_id")
	_barData15m.BeginTime = field.NewInt64(tableName, "begin_time")
	_barData15m.OpenPrice = field.NewFloat64(tableName, "open_price")
	_barData15m.ClosePrice = field.NewFloat64(tableName, "close_price")
	_barData15m.HighPrice = field.NewFloat64(tableName, "high_price")
	_barData15m.LowPrice = field.NewFloat64(tableName, "low_price")
	_barData15m.BarType = field.NewString(tableName, "bar_type")
	_barData15m.SigType = field.NewInt64(tableName, "sig_type")
	_barData15m.PriceChange = field.NewFloat64(tableName, "price_change")
	_barData15m.PriceRange = field.NewFloat64(tableName, "price_range")
	_barData15m.PriceRangePrecent = field.NewFloat64(tableName, "price_range_precent")
	_barData15m.TriggerTime = field.NewInt64(tableName, "trigger_time")

	_barData15m.fillFieldMap()

	return _barData15m
}

// barData15m 时间粒度表-每15分钟
type barData15m struct {
	barData15mDo

	ALL               field.Asterisk
	ID                field.Int64   // 主键ID
	CurrencyID        field.String  // 币种ID
	BeginTime         field.Int64   // 开始时间
	OpenPrice         field.Float64 // 开盘价格
	ClosePrice        field.Float64 // 收盘价格
	HighPrice         field.Float64 // 最高价格
	LowPrice          field.Float64 // 最低价格
	BarType           field.String  // 时间粒度类型:默认值1m,1m/5m/15m/30m/1H/4H/12H/1D/1W/1M
	SigType           field.Int64   // 行情信号: 0-/, 1-上涨, 2-快速上涨, 3-爆涨, 4-下跌, 5-快速下跌, 6-爆跌
	PriceChange       field.Float64 // 涨跌幅度
	PriceRange        field.Float64 // 涨跌振幅
	PriceRangePrecent field.Float64 // 涨跌振幅比例
	TriggerTime       field.Int64   // 触发时间

	fieldMap map[string]field.Expr
}

func (b barData15m) Table(newTableName string) *barData15m {
	b.barData15mDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b barData15m) As(alias string) *barData15m {
	b.barData15mDo.DO = *(b.barData15mDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *barData15m) updateTableName(table string) *barData15m {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.CurrencyID = field.NewString(table, "currency_id")
	b.BeginTime = field.NewInt64(table, "begin_time")
	b.OpenPrice = field.NewFloat64(table, "open_price")
	b.ClosePrice = field.NewFloat64(table, "close_price")
	b.HighPrice = field.NewFloat64(table, "high_price")
	b.LowPrice = field.NewFloat64(table, "low_price")
	b.BarType = field.NewString(table, "bar_type")
	b.SigType = field.NewInt64(table, "sig_type")
	b.PriceChange = field.NewFloat64(table, "price_change")
	b.PriceRange = field.NewFloat64(table, "price_range")
	b.PriceRangePrecent = field.NewFloat64(table, "price_range_precent")
	b.TriggerTime = field.NewInt64(table, "trigger_time")

	b.fillFieldMap()

	return b
}

func (b *barData15m) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *barData15m) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 13)
	b.fieldMap["id"] = b.ID
	b.fieldMap["currency_id"] = b.CurrencyID
	b.fieldMap["begin_time"] = b.BeginTime
	b.fieldMap["open_price"] = b.OpenPrice
	b.fieldMap["close_price"] = b.ClosePrice
	b.fieldMap["high_price"] = b.HighPrice
	b.fieldMap["low_price"] = b.LowPrice
	b.fieldMap["bar_type"] = b.BarType
	b.fieldMap["sig_type"] = b.SigType
	b.fieldMap["price_change"] = b.PriceChange
	b.fieldMap["price_range"] = b.PriceRange
	b.fieldMap["price_range_precent"] = b.PriceRangePrecent
	b.fieldMap["trigger_time"] = b.TriggerTime
}

func (b barData15m) clone(db *gorm.DB) barData15m {
	b.barData15mDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b barData15m) replaceDB(db *gorm.DB) barData15m {
	b.barData15mDo.ReplaceDB(db)
	return b
}

type barData15mDo struct{ gen.DO }

type IBarData15mDo interface {
	gen.SubQuery
	Debug() IBarData15mDo
	WithContext(ctx context.Context) IBarData15mDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBarData15mDo
	WriteDB() IBarData15mDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBarData15mDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBarData15mDo
	Not(conds ...gen.Condition) IBarData15mDo
	Or(conds ...gen.Condition) IBarData15mDo
	Select(conds ...field.Expr) IBarData15mDo
	Where(conds ...gen.Condition) IBarData15mDo
	Order(conds ...field.Expr) IBarData15mDo
	Distinct(cols ...field.Expr) IBarData15mDo
	Omit(cols ...field.Expr) IBarData15mDo
	Join(table schema.Tabler, on ...field.Expr) IBarData15mDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBarData15mDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBarData15mDo
	Group(cols ...field.Expr) IBarData15mDo
	Having(conds ...gen.Condition) IBarData15mDo
	Limit(limit int) IBarData15mDo
	Offset(offset int) IBarData15mDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBarData15mDo
	Unscoped() IBarData15mDo
	Create(values ...*model.BarData15m) error
	CreateInBatches(values []*model.BarData15m, batchSize int) error
	Save(values ...*model.BarData15m) error
	First() (*model.BarData15m, error)
	Take() (*model.BarData15m, error)
	Last() (*model.BarData15m, error)
	Find() ([]*model.BarData15m, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BarData15m, err error)
	FindInBatches(result *[]*model.BarData15m, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BarData15m) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBarData15mDo
	Assign(attrs ...field.AssignExpr) IBarData15mDo
	Joins(fields ...field.RelationField) IBarData15mDo
	Preload(fields ...field.RelationField) IBarData15mDo
	FirstOrInit() (*model.BarData15m, error)
	FirstOrCreate() (*model.BarData15m, error)
	FindByPage(offset int, limit int) (result []*model.BarData15m, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBarData15mDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b barData15mDo) Debug() IBarData15mDo {
	return b.withDO(b.DO.Debug())
}

func (b barData15mDo) WithContext(ctx context.Context) IBarData15mDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b barData15mDo) ReadDB() IBarData15mDo {
	return b.Clauses(dbresolver.Read)
}

func (b barData15mDo) WriteDB() IBarData15mDo {
	return b.Clauses(dbresolver.Write)
}

func (b barData15mDo) Session(config *gorm.Session) IBarData15mDo {
	return b.withDO(b.DO.Session(config))
}

func (b barData15mDo) Clauses(conds ...clause.Expression) IBarData15mDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b barData15mDo) Returning(value interface{}, columns ...string) IBarData15mDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b barData15mDo) Not(conds ...gen.Condition) IBarData15mDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b barData15mDo) Or(conds ...gen.Condition) IBarData15mDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b barData15mDo) Select(conds ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b barData15mDo) Where(conds ...gen.Condition) IBarData15mDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b barData15mDo) Order(conds ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b barData15mDo) Distinct(cols ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b barData15mDo) Omit(cols ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b barData15mDo) Join(table schema.Tabler, on ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b barData15mDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b barData15mDo) RightJoin(table schema.Tabler, on ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b barData15mDo) Group(cols ...field.Expr) IBarData15mDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b barData15mDo) Having(conds ...gen.Condition) IBarData15mDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b barData15mDo) Limit(limit int) IBarData15mDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b barData15mDo) Offset(offset int) IBarData15mDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b barData15mDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBarData15mDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b barData15mDo) Unscoped() IBarData15mDo {
	return b.withDO(b.DO.Unscoped())
}

func (b barData15mDo) Create(values ...*model.BarData15m) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b barData15mDo) CreateInBatches(values []*model.BarData15m, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b barData15mDo) Save(values ...*model.BarData15m) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b barData15mDo) First() (*model.BarData15m, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BarData15m), nil
	}
}

func (b barData15mDo) Take() (*model.BarData15m, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BarData15m), nil
	}
}

func (b barData15mDo) Last() (*model.BarData15m, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BarData15m), nil
	}
}

func (b barData15mDo) Find() ([]*model.BarData15m, error) {
	result, err := b.DO.Find()
	return result.([]*model.BarData15m), err
}

func (b barData15mDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BarData15m, err error) {
	buf := make([]*model.BarData15m, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b barData15mDo) FindInBatches(result *[]*model.BarData15m, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b barData15mDo) Attrs(attrs ...field.AssignExpr) IBarData15mDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b barData15mDo) Assign(attrs ...field.AssignExpr) IBarData15mDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b barData15mDo) Joins(fields ...field.RelationField) IBarData15mDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b barData15mDo) Preload(fields ...field.RelationField) IBarData15mDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b barData15mDo) FirstOrInit() (*model.BarData15m, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BarData15m), nil
	}
}

func (b barData15mDo) FirstOrCreate() (*model.BarData15m, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BarData15m), nil
	}
}

func (b barData15mDo) FindByPage(offset int, limit int) (result []*model.BarData15m, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b barData15mDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b barData15mDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b barData15mDo) Delete(models ...*model.BarData15m) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *barData15mDo) withDO(do gen.Dao) *barData15mDo {
	b.DO = *do.(*gen.DO)
	return b
}
