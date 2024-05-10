// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/ums/gen/model"
)

func newUmsIntegrationChangeHistory(db *gorm.DB, opts ...gen.DOOption) umsIntegrationChangeHistory {
	_umsIntegrationChangeHistory := umsIntegrationChangeHistory{}

	_umsIntegrationChangeHistory.umsIntegrationChangeHistoryDo.UseDB(db, opts...)
	_umsIntegrationChangeHistory.umsIntegrationChangeHistoryDo.UseModel(&model.UmsIntegrationChangeHistory{})

	tableName := _umsIntegrationChangeHistory.umsIntegrationChangeHistoryDo.TableName()
	_umsIntegrationChangeHistory.ALL = field.NewAsterisk(tableName)
	_umsIntegrationChangeHistory.ID = field.NewInt64(tableName, "id")
	_umsIntegrationChangeHistory.MemberID = field.NewInt64(tableName, "member_id")
	_umsIntegrationChangeHistory.ChangeType = field.NewInt32(tableName, "change_type")
	_umsIntegrationChangeHistory.ChangeCount = field.NewInt32(tableName, "change_count")
	_umsIntegrationChangeHistory.OperateMan = field.NewString(tableName, "operate_man")
	_umsIntegrationChangeHistory.OperateNote = field.NewString(tableName, "operate_note")
	_umsIntegrationChangeHistory.SourceType = field.NewInt32(tableName, "source_type")
	_umsIntegrationChangeHistory.CreateTime = field.NewTime(tableName, "create_time")

	_umsIntegrationChangeHistory.fillFieldMap()

	return _umsIntegrationChangeHistory
}

// umsIntegrationChangeHistory 积分变化历史记录表
type umsIntegrationChangeHistory struct {
	umsIntegrationChangeHistoryDo umsIntegrationChangeHistoryDo

	ALL         field.Asterisk
	ID          field.Int64
	MemberID    field.Int64  // 会员id
	ChangeType  field.Int32  // 改变类型：0->增加；1->减少
	ChangeCount field.Int32  // 积分改变数量
	OperateMan  field.String // 操作人员
	OperateNote field.String // 操作备注
	SourceType  field.Int32  // 积分来源：0->购物；1->管理员修改
	CreateTime  field.Time   // 创建时间

	fieldMap map[string]field.Expr
}

func (u umsIntegrationChangeHistory) Table(newTableName string) *umsIntegrationChangeHistory {
	u.umsIntegrationChangeHistoryDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsIntegrationChangeHistory) As(alias string) *umsIntegrationChangeHistory {
	u.umsIntegrationChangeHistoryDo.DO = *(u.umsIntegrationChangeHistoryDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsIntegrationChangeHistory) updateTableName(table string) *umsIntegrationChangeHistory {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.MemberID = field.NewInt64(table, "member_id")
	u.ChangeType = field.NewInt32(table, "change_type")
	u.ChangeCount = field.NewInt32(table, "change_count")
	u.OperateMan = field.NewString(table, "operate_man")
	u.OperateNote = field.NewString(table, "operate_note")
	u.SourceType = field.NewInt32(table, "source_type")
	u.CreateTime = field.NewTime(table, "create_time")

	u.fillFieldMap()

	return u
}

func (u *umsIntegrationChangeHistory) WithContext(ctx context.Context) IUmsIntegrationChangeHistoryDo {
	return u.umsIntegrationChangeHistoryDo.WithContext(ctx)
}

func (u umsIntegrationChangeHistory) TableName() string {
	return u.umsIntegrationChangeHistoryDo.TableName()
}

func (u umsIntegrationChangeHistory) Alias() string { return u.umsIntegrationChangeHistoryDo.Alias() }

func (u umsIntegrationChangeHistory) Columns(cols ...field.Expr) gen.Columns {
	return u.umsIntegrationChangeHistoryDo.Columns(cols...)
}

func (u *umsIntegrationChangeHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsIntegrationChangeHistory) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 8)
	u.fieldMap["id"] = u.ID
	u.fieldMap["member_id"] = u.MemberID
	u.fieldMap["change_type"] = u.ChangeType
	u.fieldMap["change_count"] = u.ChangeCount
	u.fieldMap["operate_man"] = u.OperateMan
	u.fieldMap["operate_note"] = u.OperateNote
	u.fieldMap["source_type"] = u.SourceType
	u.fieldMap["create_time"] = u.CreateTime
}

func (u umsIntegrationChangeHistory) clone(db *gorm.DB) umsIntegrationChangeHistory {
	u.umsIntegrationChangeHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsIntegrationChangeHistory) replaceDB(db *gorm.DB) umsIntegrationChangeHistory {
	u.umsIntegrationChangeHistoryDo.ReplaceDB(db)
	return u
}

type umsIntegrationChangeHistoryDo struct{ gen.DO }

type IUmsIntegrationChangeHistoryDo interface {
	gen.SubQuery
	Debug() IUmsIntegrationChangeHistoryDo
	WithContext(ctx context.Context) IUmsIntegrationChangeHistoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsIntegrationChangeHistoryDo
	WriteDB() IUmsIntegrationChangeHistoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsIntegrationChangeHistoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsIntegrationChangeHistoryDo
	Not(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo
	Or(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo
	Select(conds ...field.Expr) IUmsIntegrationChangeHistoryDo
	Where(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo
	Order(conds ...field.Expr) IUmsIntegrationChangeHistoryDo
	Distinct(cols ...field.Expr) IUmsIntegrationChangeHistoryDo
	Omit(cols ...field.Expr) IUmsIntegrationChangeHistoryDo
	Join(table schema.Tabler, on ...field.Expr) IUmsIntegrationChangeHistoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationChangeHistoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationChangeHistoryDo
	Group(cols ...field.Expr) IUmsIntegrationChangeHistoryDo
	Having(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo
	Limit(limit int) IUmsIntegrationChangeHistoryDo
	Offset(offset int) IUmsIntegrationChangeHistoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsIntegrationChangeHistoryDo
	Unscoped() IUmsIntegrationChangeHistoryDo
	Create(values ...*model.UmsIntegrationChangeHistory) error
	CreateInBatches(values []*model.UmsIntegrationChangeHistory, batchSize int) error
	Save(values ...*model.UmsIntegrationChangeHistory) error
	First() (*model.UmsIntegrationChangeHistory, error)
	Take() (*model.UmsIntegrationChangeHistory, error)
	Last() (*model.UmsIntegrationChangeHistory, error)
	Find() ([]*model.UmsIntegrationChangeHistory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsIntegrationChangeHistory, err error)
	FindInBatches(result *[]*model.UmsIntegrationChangeHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsIntegrationChangeHistory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsIntegrationChangeHistoryDo
	Assign(attrs ...field.AssignExpr) IUmsIntegrationChangeHistoryDo
	Joins(fields ...field.RelationField) IUmsIntegrationChangeHistoryDo
	Preload(fields ...field.RelationField) IUmsIntegrationChangeHistoryDo
	FirstOrInit() (*model.UmsIntegrationChangeHistory, error)
	FirstOrCreate() (*model.UmsIntegrationChangeHistory, error)
	FindByPage(offset int, limit int) (result []*model.UmsIntegrationChangeHistory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsIntegrationChangeHistoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsIntegrationChangeHistoryDo) Debug() IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Debug())
}

func (u umsIntegrationChangeHistoryDo) WithContext(ctx context.Context) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsIntegrationChangeHistoryDo) ReadDB() IUmsIntegrationChangeHistoryDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsIntegrationChangeHistoryDo) WriteDB() IUmsIntegrationChangeHistoryDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsIntegrationChangeHistoryDo) Session(config *gorm.Session) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsIntegrationChangeHistoryDo) Clauses(conds ...clause.Expression) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsIntegrationChangeHistoryDo) Returning(value interface{}, columns ...string) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsIntegrationChangeHistoryDo) Not(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsIntegrationChangeHistoryDo) Or(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsIntegrationChangeHistoryDo) Select(conds ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsIntegrationChangeHistoryDo) Where(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsIntegrationChangeHistoryDo) Order(conds ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsIntegrationChangeHistoryDo) Distinct(cols ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsIntegrationChangeHistoryDo) Omit(cols ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsIntegrationChangeHistoryDo) Join(table schema.Tabler, on ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsIntegrationChangeHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsIntegrationChangeHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsIntegrationChangeHistoryDo) Group(cols ...field.Expr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsIntegrationChangeHistoryDo) Having(conds ...gen.Condition) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsIntegrationChangeHistoryDo) Limit(limit int) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsIntegrationChangeHistoryDo) Offset(offset int) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsIntegrationChangeHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsIntegrationChangeHistoryDo) Unscoped() IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsIntegrationChangeHistoryDo) Create(values ...*model.UmsIntegrationChangeHistory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsIntegrationChangeHistoryDo) CreateInBatches(values []*model.UmsIntegrationChangeHistory, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsIntegrationChangeHistoryDo) Save(values ...*model.UmsIntegrationChangeHistory) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsIntegrationChangeHistoryDo) First() (*model.UmsIntegrationChangeHistory, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationChangeHistory), nil
	}
}

func (u umsIntegrationChangeHistoryDo) Take() (*model.UmsIntegrationChangeHistory, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationChangeHistory), nil
	}
}

func (u umsIntegrationChangeHistoryDo) Last() (*model.UmsIntegrationChangeHistory, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationChangeHistory), nil
	}
}

func (u umsIntegrationChangeHistoryDo) Find() ([]*model.UmsIntegrationChangeHistory, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsIntegrationChangeHistory), err
}

func (u umsIntegrationChangeHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsIntegrationChangeHistory, err error) {
	buf := make([]*model.UmsIntegrationChangeHistory, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsIntegrationChangeHistoryDo) FindInBatches(result *[]*model.UmsIntegrationChangeHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsIntegrationChangeHistoryDo) Attrs(attrs ...field.AssignExpr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsIntegrationChangeHistoryDo) Assign(attrs ...field.AssignExpr) IUmsIntegrationChangeHistoryDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsIntegrationChangeHistoryDo) Joins(fields ...field.RelationField) IUmsIntegrationChangeHistoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsIntegrationChangeHistoryDo) Preload(fields ...field.RelationField) IUmsIntegrationChangeHistoryDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsIntegrationChangeHistoryDo) FirstOrInit() (*model.UmsIntegrationChangeHistory, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationChangeHistory), nil
	}
}

func (u umsIntegrationChangeHistoryDo) FirstOrCreate() (*model.UmsIntegrationChangeHistory, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsIntegrationChangeHistory), nil
	}
}

func (u umsIntegrationChangeHistoryDo) FindByPage(offset int, limit int) (result []*model.UmsIntegrationChangeHistory, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u umsIntegrationChangeHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsIntegrationChangeHistoryDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsIntegrationChangeHistoryDo) Delete(models ...*model.UmsIntegrationChangeHistory) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsIntegrationChangeHistoryDo) withDO(do gen.Dao) *umsIntegrationChangeHistoryDo {
	u.DO = *do.(*gen.DO)
	return u
}