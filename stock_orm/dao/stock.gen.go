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

	"github.com/ocean386/common/stock_orm/model"
)

func newStock(db *gorm.DB, opts ...gen.DOOption) stock {
	_stock := stock{}

	_stock.stockDo.UseDB(db, opts...)
	_stock.stockDo.UseModel(&model.Stock{})

	tableName := _stock.stockDo.TableName()
	_stock.ALL = field.NewAsterisk(tableName)
	_stock.StockCode = field.NewString(tableName, "stock_code")
	_stock.StockName = field.NewString(tableName, "stock_name")
	_stock.TotalMarketValue = field.NewFloat64(tableName, "total_market_value")
	_stock.CirculatingMarketValue = field.NewFloat64(tableName, "circulating_market_value")
	_stock.PlateType = field.NewInt64(tableName, "plate_type")
	_stock.Industry = field.NewString(tableName, "industry")
	_stock.IndustryCode = field.NewString(tableName, "industry_code")
	_stock.Exchange = field.NewInt64(tableName, "exchange")
	_stock.MarketType = field.NewInt64(tableName, "market_type")
	_stock.IncreaseRange = field.NewFloat64(tableName, "increase_range")
	_stock.IsNewlyListed = field.NewInt64(tableName, "is_newly_listed")
	_stock.IsStStock = field.NewInt64(tableName, "is_st_stock")
	_stock.IsWatchStock = field.NewInt64(tableName, "is_watch_stock")
	_stock.ListingDate = field.NewTime(tableName, "listing_date")
	_stock.CreatedAt = field.NewTime(tableName, "created_at")
	_stock.UpdatedAt = field.NewTime(tableName, "updated_at")

	_stock.fillFieldMap()

	return _stock
}

// stock 股票列表-A股
type stock struct {
	stockDo

	ALL                    field.Asterisk
	StockCode              field.String  // 股票代码
	StockName              field.String  // 股票名称
	TotalMarketValue       field.Float64 // 总市值(亿)
	CirculatingMarketValue field.Float64 // 流通市值(亿)
	PlateType              field.Int64   // 盘股类型(0-全部,1-微小盘,2-小盘,3-中盘,4-大盘)
	Industry               field.String  // 行业
	IndustryCode           field.String  // 行业代码
	Exchange               field.Int64   // 交易所(0-全部,1-深圳,2-上海,3-北京)
	MarketType             field.Int64   // 市场类别(0-全部,1-主板10%,2-创业板20%,3-科创板20%,4-北交所30%)
	IncreaseRange          field.Float64 // 涨幅范围
	IsNewlyListed          field.Int64   // 次新股(0-否 1-是) 上市时间一年以内
	IsStStock              field.Int64   // ST股票(0-否 1-是)若ST则后期清理数据
	IsWatchStock           field.Int64   // 自选股标志(0-否 1-是)
	ListingDate            field.Time    // 上市日期
	CreatedAt              field.Time    // 创建时间
	UpdatedAt              field.Time    // 更新时间

	fieldMap map[string]field.Expr
}

func (s stock) Table(newTableName string) *stock {
	s.stockDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s stock) As(alias string) *stock {
	s.stockDo.DO = *(s.stockDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *stock) updateTableName(table string) *stock {
	s.ALL = field.NewAsterisk(table)
	s.StockCode = field.NewString(table, "stock_code")
	s.StockName = field.NewString(table, "stock_name")
	s.TotalMarketValue = field.NewFloat64(table, "total_market_value")
	s.CirculatingMarketValue = field.NewFloat64(table, "circulating_market_value")
	s.PlateType = field.NewInt64(table, "plate_type")
	s.Industry = field.NewString(table, "industry")
	s.IndustryCode = field.NewString(table, "industry_code")
	s.Exchange = field.NewInt64(table, "exchange")
	s.MarketType = field.NewInt64(table, "market_type")
	s.IncreaseRange = field.NewFloat64(table, "increase_range")
	s.IsNewlyListed = field.NewInt64(table, "is_newly_listed")
	s.IsStStock = field.NewInt64(table, "is_st_stock")
	s.IsWatchStock = field.NewInt64(table, "is_watch_stock")
	s.ListingDate = field.NewTime(table, "listing_date")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *stock) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *stock) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 16)
	s.fieldMap["stock_code"] = s.StockCode
	s.fieldMap["stock_name"] = s.StockName
	s.fieldMap["total_market_value"] = s.TotalMarketValue
	s.fieldMap["circulating_market_value"] = s.CirculatingMarketValue
	s.fieldMap["plate_type"] = s.PlateType
	s.fieldMap["industry"] = s.Industry
	s.fieldMap["industry_code"] = s.IndustryCode
	s.fieldMap["exchange"] = s.Exchange
	s.fieldMap["market_type"] = s.MarketType
	s.fieldMap["increase_range"] = s.IncreaseRange
	s.fieldMap["is_newly_listed"] = s.IsNewlyListed
	s.fieldMap["is_st_stock"] = s.IsStStock
	s.fieldMap["is_watch_stock"] = s.IsWatchStock
	s.fieldMap["listing_date"] = s.ListingDate
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s stock) clone(db *gorm.DB) stock {
	s.stockDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s stock) replaceDB(db *gorm.DB) stock {
	s.stockDo.ReplaceDB(db)
	return s
}

type stockDo struct{ gen.DO }

type IStockDo interface {
	gen.SubQuery
	Debug() IStockDo
	WithContext(ctx context.Context) IStockDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStockDo
	WriteDB() IStockDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStockDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStockDo
	Not(conds ...gen.Condition) IStockDo
	Or(conds ...gen.Condition) IStockDo
	Select(conds ...field.Expr) IStockDo
	Where(conds ...gen.Condition) IStockDo
	Order(conds ...field.Expr) IStockDo
	Distinct(cols ...field.Expr) IStockDo
	Omit(cols ...field.Expr) IStockDo
	Join(table schema.Tabler, on ...field.Expr) IStockDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStockDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStockDo
	Group(cols ...field.Expr) IStockDo
	Having(conds ...gen.Condition) IStockDo
	Limit(limit int) IStockDo
	Offset(offset int) IStockDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStockDo
	Unscoped() IStockDo
	Create(values ...*model.Stock) error
	CreateInBatches(values []*model.Stock, batchSize int) error
	Save(values ...*model.Stock) error
	First() (*model.Stock, error)
	Take() (*model.Stock, error)
	Last() (*model.Stock, error)
	Find() ([]*model.Stock, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Stock, err error)
	FindInBatches(result *[]*model.Stock, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Stock) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStockDo
	Assign(attrs ...field.AssignExpr) IStockDo
	Joins(fields ...field.RelationField) IStockDo
	Preload(fields ...field.RelationField) IStockDo
	FirstOrInit() (*model.Stock, error)
	FirstOrCreate() (*model.Stock, error)
	FindByPage(offset int, limit int) (result []*model.Stock, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStockDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s stockDo) Debug() IStockDo {
	return s.withDO(s.DO.Debug())
}

func (s stockDo) WithContext(ctx context.Context) IStockDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s stockDo) ReadDB() IStockDo {
	return s.Clauses(dbresolver.Read)
}

func (s stockDo) WriteDB() IStockDo {
	return s.Clauses(dbresolver.Write)
}

func (s stockDo) Session(config *gorm.Session) IStockDo {
	return s.withDO(s.DO.Session(config))
}

func (s stockDo) Clauses(conds ...clause.Expression) IStockDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s stockDo) Returning(value interface{}, columns ...string) IStockDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s stockDo) Not(conds ...gen.Condition) IStockDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s stockDo) Or(conds ...gen.Condition) IStockDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s stockDo) Select(conds ...field.Expr) IStockDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s stockDo) Where(conds ...gen.Condition) IStockDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s stockDo) Order(conds ...field.Expr) IStockDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s stockDo) Distinct(cols ...field.Expr) IStockDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s stockDo) Omit(cols ...field.Expr) IStockDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s stockDo) Join(table schema.Tabler, on ...field.Expr) IStockDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s stockDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStockDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s stockDo) RightJoin(table schema.Tabler, on ...field.Expr) IStockDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s stockDo) Group(cols ...field.Expr) IStockDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s stockDo) Having(conds ...gen.Condition) IStockDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s stockDo) Limit(limit int) IStockDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s stockDo) Offset(offset int) IStockDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s stockDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStockDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s stockDo) Unscoped() IStockDo {
	return s.withDO(s.DO.Unscoped())
}

func (s stockDo) Create(values ...*model.Stock) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s stockDo) CreateInBatches(values []*model.Stock, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s stockDo) Save(values ...*model.Stock) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s stockDo) First() (*model.Stock, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Stock), nil
	}
}

func (s stockDo) Take() (*model.Stock, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Stock), nil
	}
}

func (s stockDo) Last() (*model.Stock, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Stock), nil
	}
}

func (s stockDo) Find() ([]*model.Stock, error) {
	result, err := s.DO.Find()
	return result.([]*model.Stock), err
}

func (s stockDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Stock, err error) {
	buf := make([]*model.Stock, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s stockDo) FindInBatches(result *[]*model.Stock, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s stockDo) Attrs(attrs ...field.AssignExpr) IStockDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s stockDo) Assign(attrs ...field.AssignExpr) IStockDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s stockDo) Joins(fields ...field.RelationField) IStockDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s stockDo) Preload(fields ...field.RelationField) IStockDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s stockDo) FirstOrInit() (*model.Stock, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Stock), nil
	}
}

func (s stockDo) FirstOrCreate() (*model.Stock, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Stock), nil
	}
}

func (s stockDo) FindByPage(offset int, limit int) (result []*model.Stock, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s stockDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s stockDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s stockDo) Delete(models ...*model.Stock) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *stockDo) withDO(do gen.Dao) *stockDo {
	s.DO = *do.(*gen.DO)
	return s
}
