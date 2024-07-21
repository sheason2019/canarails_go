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

func newApp(db *gorm.DB, opts ...gen.DOOption) app {
	_app := app{}

	_app.appDo.UseDB(db, opts...)
	_app.appDo.UseModel(&models.App{})

	tableName := _app.appDo.TableName()
	_app.ALL = field.NewAsterisk(tableName)
	_app.ID = field.NewUint(tableName, "id")
	_app.CreatedAt = field.NewTime(tableName, "created_at")
	_app.UpdatedAt = field.NewTime(tableName, "updated_at")
	_app.DeletedAt = field.NewField(tableName, "deleted_at")
	_app.Title = field.NewString(tableName, "title")
	_app.Description = field.NewString(tableName, "description")
	_app.Hostnames = field.NewField(tableName, "hostnames")
	_app.AppVariants = appHasManyAppVariants{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AppVariants", "models.AppVariant"),
		App: struct {
			field.RelationField
			AppVariants struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("AppVariants.App", "models.App"),
			AppVariants: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("AppVariants.App.AppVariants", "models.AppVariant"),
			},
		},
	}

	_app.fillFieldMap()

	return _app
}

type app struct {
	appDo appDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Title       field.String
	Description field.String
	Hostnames   field.Field
	AppVariants appHasManyAppVariants

	fieldMap map[string]field.Expr
}

func (a app) Table(newTableName string) *app {
	a.appDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a app) As(alias string) *app {
	a.appDo.DO = *(a.appDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *app) updateTableName(table string) *app {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Title = field.NewString(table, "title")
	a.Description = field.NewString(table, "description")
	a.Hostnames = field.NewField(table, "hostnames")

	a.fillFieldMap()

	return a
}

func (a *app) WithContext(ctx context.Context) IAppDo { return a.appDo.WithContext(ctx) }

func (a app) TableName() string { return a.appDo.TableName() }

func (a app) Alias() string { return a.appDo.Alias() }

func (a app) Columns(cols ...field.Expr) gen.Columns { return a.appDo.Columns(cols...) }

func (a *app) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *app) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["title"] = a.Title
	a.fieldMap["description"] = a.Description
	a.fieldMap["hostnames"] = a.Hostnames

}

func (a app) clone(db *gorm.DB) app {
	a.appDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a app) replaceDB(db *gorm.DB) app {
	a.appDo.ReplaceDB(db)
	return a
}

type appHasManyAppVariants struct {
	db *gorm.DB

	field.RelationField

	App struct {
		field.RelationField
		AppVariants struct {
			field.RelationField
		}
	}
}

func (a appHasManyAppVariants) Where(conds ...field.Expr) *appHasManyAppVariants {
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

func (a appHasManyAppVariants) WithContext(ctx context.Context) *appHasManyAppVariants {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a appHasManyAppVariants) Session(session *gorm.Session) *appHasManyAppVariants {
	a.db = a.db.Session(session)
	return &a
}

func (a appHasManyAppVariants) Model(m *models.App) *appHasManyAppVariantsTx {
	return &appHasManyAppVariantsTx{a.db.Model(m).Association(a.Name())}
}

type appHasManyAppVariantsTx struct{ tx *gorm.Association }

func (a appHasManyAppVariantsTx) Find() (result []*models.AppVariant, err error) {
	return result, a.tx.Find(&result)
}

func (a appHasManyAppVariantsTx) Append(values ...*models.AppVariant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a appHasManyAppVariantsTx) Replace(values ...*models.AppVariant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a appHasManyAppVariantsTx) Delete(values ...*models.AppVariant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a appHasManyAppVariantsTx) Clear() error {
	return a.tx.Clear()
}

func (a appHasManyAppVariantsTx) Count() int64 {
	return a.tx.Count()
}

type appDo struct{ gen.DO }

type IAppDo interface {
	gen.SubQuery
	Debug() IAppDo
	WithContext(ctx context.Context) IAppDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAppDo
	WriteDB() IAppDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAppDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppDo
	Not(conds ...gen.Condition) IAppDo
	Or(conds ...gen.Condition) IAppDo
	Select(conds ...field.Expr) IAppDo
	Where(conds ...gen.Condition) IAppDo
	Order(conds ...field.Expr) IAppDo
	Distinct(cols ...field.Expr) IAppDo
	Omit(cols ...field.Expr) IAppDo
	Join(table schema.Tabler, on ...field.Expr) IAppDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppDo
	Group(cols ...field.Expr) IAppDo
	Having(conds ...gen.Condition) IAppDo
	Limit(limit int) IAppDo
	Offset(offset int) IAppDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppDo
	Unscoped() IAppDo
	Create(values ...*models.App) error
	CreateInBatches(values []*models.App, batchSize int) error
	Save(values ...*models.App) error
	First() (*models.App, error)
	Take() (*models.App, error)
	Last() (*models.App, error)
	Find() ([]*models.App, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.App, err error)
	FindInBatches(result *[]*models.App, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.App) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppDo
	Assign(attrs ...field.AssignExpr) IAppDo
	Joins(fields ...field.RelationField) IAppDo
	Preload(fields ...field.RelationField) IAppDo
	FirstOrInit() (*models.App, error)
	FirstOrCreate() (*models.App, error)
	FindByPage(offset int, limit int) (result []*models.App, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a appDo) Debug() IAppDo {
	return a.withDO(a.DO.Debug())
}

func (a appDo) WithContext(ctx context.Context) IAppDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appDo) ReadDB() IAppDo {
	return a.Clauses(dbresolver.Read)
}

func (a appDo) WriteDB() IAppDo {
	return a.Clauses(dbresolver.Write)
}

func (a appDo) Session(config *gorm.Session) IAppDo {
	return a.withDO(a.DO.Session(config))
}

func (a appDo) Clauses(conds ...clause.Expression) IAppDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appDo) Returning(value interface{}, columns ...string) IAppDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appDo) Not(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appDo) Or(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appDo) Select(conds ...field.Expr) IAppDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appDo) Where(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appDo) Order(conds ...field.Expr) IAppDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appDo) Distinct(cols ...field.Expr) IAppDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appDo) Omit(cols ...field.Expr) IAppDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appDo) Join(table schema.Tabler, on ...field.Expr) IAppDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appDo) Group(cols ...field.Expr) IAppDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appDo) Having(conds ...gen.Condition) IAppDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appDo) Limit(limit int) IAppDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appDo) Offset(offset int) IAppDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appDo) Unscoped() IAppDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appDo) Create(values ...*models.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appDo) CreateInBatches(values []*models.App, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appDo) Save(values ...*models.App) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appDo) First() (*models.App, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.App), nil
	}
}

func (a appDo) Take() (*models.App, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.App), nil
	}
}

func (a appDo) Last() (*models.App, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.App), nil
	}
}

func (a appDo) Find() ([]*models.App, error) {
	result, err := a.DO.Find()
	return result.([]*models.App), err
}

func (a appDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.App, err error) {
	buf := make([]*models.App, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appDo) FindInBatches(result *[]*models.App, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appDo) Attrs(attrs ...field.AssignExpr) IAppDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appDo) Assign(attrs ...field.AssignExpr) IAppDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appDo) Joins(fields ...field.RelationField) IAppDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appDo) Preload(fields ...field.RelationField) IAppDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appDo) FirstOrInit() (*models.App, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.App), nil
	}
}

func (a appDo) FirstOrCreate() (*models.App, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.App), nil
	}
}

func (a appDo) FindByPage(offset int, limit int) (result []*models.App, count int64, err error) {
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

func (a appDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appDo) Delete(models ...*models.App) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appDo) withDO(do gen.Dao) *appDo {
	a.DO = *do.(*gen.DO)
	return a
}
