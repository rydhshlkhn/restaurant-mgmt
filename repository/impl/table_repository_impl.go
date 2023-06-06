package impl

import (
	"context"
	"strings"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type tableRepositoryImpl struct {
	DB *gorm.DB
}

func NewTableRepositoryImpl(DB *gorm.DB) repository.TableRepository {
	return &tableRepositoryImpl{DB}
}

func (r *tableRepositoryImpl) FetchAll(ctx context.Context, filter *model.TableFilter) (tables []entity.Table, err error) {
	err = r.setFilter(r.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&tables).Error
	return
}

func (r *tableRepositoryImpl) Count(ctx context.Context, filter *model.TableFilter) (count int) {
	var total int64
	r.setFilter(r.DB, filter).Model(&entity.Table{}).Count(&total)
	count = int(total)
	return
}

func (r *tableRepositoryImpl) Find(ctx context.Context, filter *model.TableFilter) (result entity.Table, err error) {
	err = r.setFilter(r.DB, filter).First(&result).Error
	return
}

func (r *tableRepositoryImpl) Save(ctx context.Context, table *entity.Table) (*entity.Table, error) {
	if table.ID == "" {
		table.ID = uuid.NewString()
		err := r.DB.Create(table).Error
		return table, err
	}
	err := r.DB.Save(table).Error
	return table, err
}

func (r *tableRepositoryImpl) setFilter(db *gorm.DB, filter *model.TableFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	return db
}
