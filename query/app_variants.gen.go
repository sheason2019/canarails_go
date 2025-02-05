// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"canarails.dev/database/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newAppVariant(db *gorm.DB, opts ...gen.DOOption) appVariant {
	_appVariant := appVariant{}

	_appVariant.appVariantDo.UseDB(db, opts...)
	_appVariant.appVariantDo.UseModel(&models.AppVariant{})

	tableName := _appVariant.appVariantDo.TableName()
	_appVariant.ALL = field.NewAsterisk(tableName)
	_appVariant.ID = field.NewUint(tableName, "id")
	_appVariant.CreatedAt = field.NewTime(tableName, "created_at")
	_appVariant.UpdatedAt = field.NewTime(tableName, "updated_at")
	_appVariant.DeletedAt = field.NewField(tableName, "deleted_at")
	_appVariant.Title = field.NewString(tableName, "title")
	_appVariant.Description = field.NewString(tableName, "description")
	_appVariant.ExposePort = field.NewUint(tableName, "expose_port")
	_appVariant.Matches = field.NewField(tableName, "matches")
	_appVariant.ImageName = field.NewString(tableName, "image_name")
	_appVariant.Replicas = field.NewUint(tableName, "replicas")
	_appVariant.AppID = field.NewUint(tableName, "app_id")
	_appVariant.App = appVariantBelongsToApp{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("App", "models.App"),
		AppVariants: struct {
			field.RelationField
			App struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("App.AppVariants", "models.AppVariant"),
			App: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("App.AppVariants.App", "models.App"),
			},
		},
	}

	_appVariant.fillFieldMap()

	return _appVariant
}

type appVariant struct {
	appVariantDo appVariantDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Title       field.String
	Description field.String
	ExposePort  field.Uint
	Matches     field.Field
	ImageName   field.String
	Replicas    field.Uint
	AppID       field.Uint
	App         appVariantBelongsToApp

	fieldMap map[string]field.Expr
}

func (a appVariant) Table(newTableName string) *appVariant {
	a.appVariantDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appVariant) As(alias string) *appVariant {
	a.appVariantDo.DO = *(a.appVariantDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appVariant) updateTableName(table string) *appVariant {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Title = field.NewString(table, "title")
	a.Description = field.NewString(table, "description")
	a.ExposePort = field.NewUint(table, "expose_port")
	a.Matches = field.NewField(table, "matches")
	a.ImageName = field.NewString(table, "image_name")
	a.Replicas = field.NewUint(table, "replicas")
	a.AppID = field.NewUint(table, "app_id")

	a.fillFieldMap()

	return a
}

func (a *appVariant) WithContext(ctx context.Context) IAppVariantDo {
	return a.appVariantDo.WithContext(ctx)
}

func (a appVariant) TableName() string { return a.appVariantDo.TableName() }

func (a appVariant) Alias() string { return a.appVariantDo.Alias() }

func (a appVariant) Columns(cols ...field.Expr) gen.Columns { return a.appVariantDo.Columns(cols...) }

func (a *appVariant) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appVariant) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 12)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["title"] = a.Title
	a.fieldMap["description"] = a.Description
	a.fieldMap["expose_port"] = a.ExposePort
	a.fieldMap["matches"] = a.Matches
	a.fieldMap["image_name"] = a.ImageName
	a.fieldMap["replicas"] = a.Replicas
	a.fieldMap["app_id"] = a.AppID

}

func (a appVariant) clone(db *gorm.DB) appVariant {
	a.appVariantDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a appVariant) replaceDB(db *gorm.DB) appVariant {
	a.appVariantDo.ReplaceDB(db)
	return a
}

type appVariantBelongsToApp struct {
	db *gorm.DB

	field.RelationField

	AppVariants struct {
		field.RelationField
		App struct {
			field.RelationField
		}
	}
}

func (a appVariantBelongsToApp) Where(conds ...field.Expr) *appVariantBelongsToApp {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a appVariantBelongsToApp) WithContext(ctx context.Context) *appVariantBelongsToApp {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a appVariantBelongsToApp) Session(session *gorm.Session) *appVariantBelongsToApp {
	a.db = a.db.Session(session)
	return &a
}

func (a appVariantBelongsToApp) Model(m *models.AppVariant) *appVariantBelongsToAppTx {
	return &appVariantBelongsToAppTx{a.db.Model(m).Association(a.Name())}
}

type appVariantBelongsToAppTx struct{ tx *gorm.Association }

func (a appVariantBelongsToAppTx) Find() (result *models.App, err error) {
	return result, a.tx.Find(&result)
}

func (a appVariantBelongsToAppTx) Append(values ...*models.App) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a appVariantBelongsToAppTx) Replace(values ...*models.App) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a appVariantBelongsToAppTx) Delete(values ...*models.App) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a appVariantBelongsToAppTx) Clear() error {
	return a.tx.Clear()
}

func (a appVariantBelongsToAppTx) Count() int64 {
	return a.tx.Count()
}

type appVariantDo struct{ gen.DO }

type IAppVariantDo interface {
	gen.SubQuery
	Debug() IAppVariantDo
	WithContext(ctx context.Context) IAppVariantDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAppVariantDo
	WriteDB() IAppVariantDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAppVariantDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppVariantDo
	Not(conds ...gen.Condition) IAppVariantDo
	Or(conds ...gen.Condition) IAppVariantDo
	Select(conds ...field.Expr) IAppVariantDo
	Where(conds ...gen.Condition) IAppVariantDo
	Order(conds ...field.Expr) IAppVariantDo
	Distinct(cols ...field.Expr) IAppVariantDo
	Omit(cols ...field.Expr) IAppVariantDo
	Join(table schema.Tabler, on ...field.Expr) IAppVariantDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppVariantDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppVariantDo
	Group(cols ...field.Expr) IAppVariantDo
	Having(conds ...gen.Condition) IAppVariantDo
	Limit(limit int) IAppVariantDo
	Offset(offset int) IAppVariantDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppVariantDo
	Unscoped() IAppVariantDo
	Create(values ...*models.AppVariant) error
	CreateInBatches(values []*models.AppVariant, batchSize int) error
	Save(values ...*models.AppVariant) error
	First() (*models.AppVariant, error)
	Take() (*models.AppVariant, error)
	Last() (*models.AppVariant, error)
	Find() ([]*models.AppVariant, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AppVariant, err error)
	FindInBatches(result *[]*models.AppVariant, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.AppVariant) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppVariantDo
	Assign(attrs ...field.AssignExpr) IAppVariantDo
	Joins(fields ...field.RelationField) IAppVariantDo
	Preload(fields ...field.RelationField) IAppVariantDo
	FirstOrInit() (*models.AppVariant, error)
	FirstOrCreate() (*models.AppVariant, error)
	FindByPage(offset int, limit int) (result []*models.AppVariant, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppVariantDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a appVariantDo) Debug() IAppVariantDo {
	return a.withDO(a.DO.Debug())
}

func (a appVariantDo) WithContext(ctx context.Context) IAppVariantDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appVariantDo) ReadDB() IAppVariantDo {
	return a.Clauses(dbresolver.Read)
}

func (a appVariantDo) WriteDB() IAppVariantDo {
	return a.Clauses(dbresolver.Write)
}

func (a appVariantDo) Session(config *gorm.Session) IAppVariantDo {
	return a.withDO(a.DO.Session(config))
}

func (a appVariantDo) Clauses(conds ...clause.Expression) IAppVariantDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appVariantDo) Returning(value interface{}, columns ...string) IAppVariantDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appVariantDo) Not(conds ...gen.Condition) IAppVariantDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appVariantDo) Or(conds ...gen.Condition) IAppVariantDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appVariantDo) Select(conds ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appVariantDo) Where(conds ...gen.Condition) IAppVariantDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appVariantDo) Order(conds ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appVariantDo) Distinct(cols ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appVariantDo) Omit(cols ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appVariantDo) Join(table schema.Tabler, on ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appVariantDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appVariantDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appVariantDo) Group(cols ...field.Expr) IAppVariantDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appVariantDo) Having(conds ...gen.Condition) IAppVariantDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appVariantDo) Limit(limit int) IAppVariantDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appVariantDo) Offset(offset int) IAppVariantDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appVariantDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppVariantDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appVariantDo) Unscoped() IAppVariantDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appVariantDo) Create(values ...*models.AppVariant) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appVariantDo) CreateInBatches(values []*models.AppVariant, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appVariantDo) Save(values ...*models.AppVariant) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appVariantDo) First() (*models.AppVariant, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.AppVariant), nil
	}
}

func (a appVariantDo) Take() (*models.AppVariant, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.AppVariant), nil
	}
}

func (a appVariantDo) Last() (*models.AppVariant, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.AppVariant), nil
	}
}

func (a appVariantDo) Find() ([]*models.AppVariant, error) {
	result, err := a.DO.Find()
	return result.([]*models.AppVariant), err
}

func (a appVariantDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AppVariant, err error) {
	buf := make([]*models.AppVariant, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appVariantDo) FindInBatches(result *[]*models.AppVariant, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appVariantDo) Attrs(attrs ...field.AssignExpr) IAppVariantDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appVariantDo) Assign(attrs ...field.AssignExpr) IAppVariantDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appVariantDo) Joins(fields ...field.RelationField) IAppVariantDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appVariantDo) Preload(fields ...field.RelationField) IAppVariantDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appVariantDo) FirstOrInit() (*models.AppVariant, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.AppVariant), nil
	}
}

func (a appVariantDo) FirstOrCreate() (*models.AppVariant, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.AppVariant), nil
	}
}

func (a appVariantDo) FindByPage(offset int, limit int) (result []*models.AppVariant, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appVariantDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appVariantDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appVariantDo) Delete(models ...*models.AppVariant) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appVariantDo) withDO(do gen.Dao) *appVariantDo {
	a.DO = *do.(*gen.DO)
	return a
}
