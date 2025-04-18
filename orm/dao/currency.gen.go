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

func newCurrency(db *gorm.DB, opts ...gen.DOOption) currency {
	_currency := currency{}

	_currency.currencyDo.UseDB(db, opts...)
	_currency.currencyDo.UseModel(&model.Currency{})

	tableName := _currency.currencyDo.TableName()
	_currency.ALL = field.NewAsterisk(tableName)
	_currency.ID = field.NewInt64(tableName, "id")
	_currency.CurrencyID = field.NewString(tableName, "currency_id")
	_currency.CurrencyName = field.NewString(tableName, "currency_name")
	_currency.MinUnitPrice = field.NewFloat64(tableName, "min_unit_price")
	_currency.CreatedAt = field.NewInt64(tableName, "created_at")
	_currency.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_currency.Sort = field.NewInt64(tableName, "sort")

	_currency.fillFieldMap()

	return _currency
}

// currency 虚拟币币种表
type currency struct {
	currencyDo

	ALL          field.Asterisk
	ID           field.Int64   // 主键ID
	CurrencyID   field.String  // 币种ID
	CurrencyName field.String  // 币种-名称
	MinUnitPrice field.Float64 // 最小单元-价格
	CreatedAt    field.Int64   // 创建时间
	UpdatedAt    field.Int64   // 更新时间
	Sort         field.Int64   // 序号

	fieldMap map[string]field.Expr
}

func (c currency) Table(newTableName string) *currency {
	c.currencyDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c currency) As(alias string) *currency {
	c.currencyDo.DO = *(c.currencyDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *currency) updateTableName(table string) *currency {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CurrencyID = field.NewString(table, "currency_id")
	c.CurrencyName = field.NewString(table, "currency_name")
	c.MinUnitPrice = field.NewFloat64(table, "min_unit_price")
	c.CreatedAt = field.NewInt64(table, "created_at")
	c.UpdatedAt = field.NewInt64(table, "updated_at")
	c.Sort = field.NewInt64(table, "sort")

	c.fillFieldMap()

	return c
}

func (c *currency) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *currency) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["currency_id"] = c.CurrencyID
	c.fieldMap["currency_name"] = c.CurrencyName
	c.fieldMap["min_unit_price"] = c.MinUnitPrice
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["sort"] = c.Sort
}

func (c currency) clone(db *gorm.DB) currency {
	c.currencyDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c currency) replaceDB(db *gorm.DB) currency {
	c.currencyDo.ReplaceDB(db)
	return c
}

type currencyDo struct{ gen.DO }

type ICurrencyDo interface {
	gen.SubQuery
	Debug() ICurrencyDo
	WithContext(ctx context.Context) ICurrencyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurrencyDo
	WriteDB() ICurrencyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurrencyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurrencyDo
	Not(conds ...gen.Condition) ICurrencyDo
	Or(conds ...gen.Condition) ICurrencyDo
	Select(conds ...field.Expr) ICurrencyDo
	Where(conds ...gen.Condition) ICurrencyDo
	Order(conds ...field.Expr) ICurrencyDo
	Distinct(cols ...field.Expr) ICurrencyDo
	Omit(cols ...field.Expr) ICurrencyDo
	Join(table schema.Tabler, on ...field.Expr) ICurrencyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurrencyDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurrencyDo
	Group(cols ...field.Expr) ICurrencyDo
	Having(conds ...gen.Condition) ICurrencyDo
	Limit(limit int) ICurrencyDo
	Offset(offset int) ICurrencyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurrencyDo
	Unscoped() ICurrencyDo
	Create(values ...*model.Currency) error
	CreateInBatches(values []*model.Currency, batchSize int) error
	Save(values ...*model.Currency) error
	First() (*model.Currency, error)
	Take() (*model.Currency, error)
	Last() (*model.Currency, error)
	Find() ([]*model.Currency, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Currency, err error)
	FindInBatches(result *[]*model.Currency, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Currency) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurrencyDo
	Assign(attrs ...field.AssignExpr) ICurrencyDo
	Joins(fields ...field.RelationField) ICurrencyDo
	Preload(fields ...field.RelationField) ICurrencyDo
	FirstOrInit() (*model.Currency, error)
	FirstOrCreate() (*model.Currency, error)
	FindByPage(offset int, limit int) (result []*model.Currency, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurrencyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c currencyDo) Debug() ICurrencyDo {
	return c.withDO(c.DO.Debug())
}

func (c currencyDo) WithContext(ctx context.Context) ICurrencyDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c currencyDo) ReadDB() ICurrencyDo {
	return c.Clauses(dbresolver.Read)
}

func (c currencyDo) WriteDB() ICurrencyDo {
	return c.Clauses(dbresolver.Write)
}

func (c currencyDo) Session(config *gorm.Session) ICurrencyDo {
	return c.withDO(c.DO.Session(config))
}

func (c currencyDo) Clauses(conds ...clause.Expression) ICurrencyDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c currencyDo) Returning(value interface{}, columns ...string) ICurrencyDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c currencyDo) Not(conds ...gen.Condition) ICurrencyDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c currencyDo) Or(conds ...gen.Condition) ICurrencyDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c currencyDo) Select(conds ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c currencyDo) Where(conds ...gen.Condition) ICurrencyDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c currencyDo) Order(conds ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c currencyDo) Distinct(cols ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c currencyDo) Omit(cols ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c currencyDo) Join(table schema.Tabler, on ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c currencyDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c currencyDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c currencyDo) Group(cols ...field.Expr) ICurrencyDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c currencyDo) Having(conds ...gen.Condition) ICurrencyDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c currencyDo) Limit(limit int) ICurrencyDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c currencyDo) Offset(offset int) ICurrencyDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c currencyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurrencyDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c currencyDo) Unscoped() ICurrencyDo {
	return c.withDO(c.DO.Unscoped())
}

func (c currencyDo) Create(values ...*model.Currency) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c currencyDo) CreateInBatches(values []*model.Currency, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c currencyDo) Save(values ...*model.Currency) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c currencyDo) First() (*model.Currency, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Currency), nil
	}
}

func (c currencyDo) Take() (*model.Currency, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Currency), nil
	}
}

func (c currencyDo) Last() (*model.Currency, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Currency), nil
	}
}

func (c currencyDo) Find() ([]*model.Currency, error) {
	result, err := c.DO.Find()
	return result.([]*model.Currency), err
}

func (c currencyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Currency, err error) {
	buf := make([]*model.Currency, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c currencyDo) FindInBatches(result *[]*model.Currency, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c currencyDo) Attrs(attrs ...field.AssignExpr) ICurrencyDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c currencyDo) Assign(attrs ...field.AssignExpr) ICurrencyDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c currencyDo) Joins(fields ...field.RelationField) ICurrencyDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c currencyDo) Preload(fields ...field.RelationField) ICurrencyDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c currencyDo) FirstOrInit() (*model.Currency, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Currency), nil
	}
}

func (c currencyDo) FirstOrCreate() (*model.Currency, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Currency), nil
	}
}

func (c currencyDo) FindByPage(offset int, limit int) (result []*model.Currency, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c currencyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c currencyDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c currencyDo) Delete(models ...*model.Currency) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *currencyDo) withDO(do gen.Dao) *currencyDo {
	c.DO = *do.(*gen.DO)
	return c
}
