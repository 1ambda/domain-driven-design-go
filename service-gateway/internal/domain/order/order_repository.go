package order

import (
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	AddOrder(order *Order) (*Order, e.Exception)
	FindOrderById(id uint) (*Order, e.Exception)
	AddOrderDetail(order *OrderDetail) (*OrderDetail, e.Exception)
	FindOrderDetailById(id uint) (*OrderDetail, e.Exception)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) AddOrder(record *Order) (*Order, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		return nil, e.NewInternalServerException(err, "Failed to create Order")
	}

	return record, nil
}

func (r *repositoryImpl) FindOrderById(id uint) (*Order, e.Exception) {
	record := &Order{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(err, "Failed to find Order does not exist")
		}

		return nil, e.NewInternalServerException(err, "Failed to find Order")
	}

	return record, nil
}

func (r *repositoryImpl) AddOrderDetail(record *OrderDetail) (*OrderDetail, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		return nil, e.NewInternalServerException(err, "Failed to create OrderDetail")
	}

	return record, nil
}

func (r *repositoryImpl) FindOrderDetailById(id uint) (*OrderDetail, e.Exception) {
	record := &OrderDetail{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(err, "Failed to find OrderDetail does not exist")
		}

		return nil, e.NewInternalServerException(err, "Failed to find OrderDEtail")
	}

	return record, nil
}
