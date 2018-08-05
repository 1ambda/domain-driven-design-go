package product

import (
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	AddCategory(record *Category) (*Category, e.Exception)
	FindCategoryById(id uint) (*Category, e.Exception)

	AddImage(record *Image) (*Image, e.Exception)
	FindImageById(id uint) (*Image, e.Exception)

	AddProduct(record *Product) (*Product, e.Exception)
	FindProductWithOptions(id uint) (*Product, []*ProductOption, e.Exception)
	FindAllProducts(itemCountPerPage int, currentPageOffset int) (int, []*Product, e.Exception)
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

type repositoryImpl struct {
	db *gorm.DB
}

func (r *repositoryImpl) AddCategory(record *Category) (*Category, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create Category")
		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) AddImage(record *Image) (*Image, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create Image")
		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FindCategoryById(id uint) (*Category, e.Exception) {
	record := &Category{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find Category by id")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FindImageById(id uint) (*Image, e.Exception) {
	record := &Image{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find Image by id")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) AddProduct(record *Product) (*Product, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create Product")
		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FindProductWithOptions(id uint) (*Product, []*ProductOption, e.Exception) {
	record := &Product{}

	tx := r.db.Begin()
	err := tx.Where("id = ?", id).
		Preload("Category").
		Preload("Image").
		First(record).
		Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find Product by id")
		tx.Rollback()

		if gorm.IsRecordNotFoundError(err) {
			return nil, nil, e.NewNotFoundException(wrap)
		}
		return nil, nil, e.NewInternalServerException(wrap)
	}

	var productOptions []*ProductOption

	if err := tx.
		Where("product_id = ?", id).
		Find(&productOptions).
		Error; err != nil {

		wrap := errors.Wrap(err, "Failed to get ProductOption list")
		tx.Rollback()
		return nil, nil, e.NewInternalServerException(wrap)
	}

	tx.Commit()

	return record, productOptions, nil
}

func (r *repositoryImpl) FindAllProducts(itemCountPerPage int, currentPageOffset int) (int, []*Product, e.Exception) {
	tx := r.db.Begin()

	product := &Product{}
	var productList []*Product
	count := 0

	dbOffset := currentPageOffset * itemCountPerPage
	dbLimit := itemCountPerPage

	if err := tx.
		Table(product.TableName()).
		Count(&count).
		Preload("Category").
		Preload("Image").
		Order("created_at asc").
		Offset(dbOffset).
		Limit(dbLimit).
		Find(&productList).
		Error; err != nil {

		tx.Rollback()
		wrap := errors.Wrap(err, "Failed to get Product list")
		return 0, nil, e.NewInternalServerException(wrap)
	}

	tx.Commit()
	return count, productList, nil
}
