package impl

import (
	"context"
	"strings"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepositoryImpl(DB *gorm.DB) repository.OrderRepository {
	return &orderRepositoryImpl{DB}
}

func (r *orderRepositoryImpl) FetchAll(ctx context.Context, filter *model.OrderFilter) (orders []entity.Order, err error) {
	err = r.setFilter(r.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&orders).Error
	return
}

func (r *orderRepositoryImpl) Count(ctx context.Context, filter *model.OrderFilter) (count int) {
	var total int64
	r.setFilter(r.DB, filter).Model(&entity.Order{}).Count(&total)
	count = int(total)
	return
}

func (r *orderRepositoryImpl) Find(ctx context.Context, filter *model.OrderFilter) (result *entity.Order, err error) {
	err = r.setFilter(r.DB, filter).First(&result).Error
	return
}

func (r *orderRepositoryImpl) Save(tx *gorm.DB, order *entity.Order) (result *entity.Order, err error) {
	if order.ID == "" {
		order.ID = uuid.NewString()
		err = tx.Create(order).Error
	} else {
		err = tx.Save(order).Error
	}

	if err != nil {return}

	err = tx.Preload("Table").First(&result, "id = ?", order.ID).Error
	return

}

func (r *orderRepositoryImpl) setFilter(db *gorm.DB, filter *model.OrderFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	for _, preload := range(filter.Preloads) {
		db = db.Preload(preload)
	}

	return db
}