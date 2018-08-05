package user

import (
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	DeleteUser(id uint) (bool, e.Exception)
	FindUserById(id uint) (*User, e.Exception)
	FineAllUsers() (*[]User, e.Exception)

	CreateAuthIdentity(uid string, email string, password string) (*AuthIdentity, e.Exception)
	FindAuthIdentityByUID(uid string) (*AuthIdentity, e.Exception)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) DeleteUser(id uint) (bool, e.Exception) {
	record := &User{}
	result := r.db.Where("id = ?", id).Delete(record)

	if result.Error != nil {
		wrap := errors.Wrap(result.Error, "Failed to delete User")
		return false, e.NewInternalServerException(wrap)
	}

	if result.RowsAffected < 1 {
		wrap := errors.Wrap(result.Error, "Failed to fine User to be deleted")
		return false, e.NewNotFoundException(wrap)
	}

	return true, nil
}

func (r *repositoryImpl) FindUserById(id uint) (*User, e.Exception) {
	record := &User{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find User")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FineAllUsers() (*[]User, e.Exception) {
	// TODO: use db.tx

	var records []User

	err := r.db.Find(&records).Error
	if err != nil {
		wrap := errors.Wrap(err, "Failed to find all User")
		return nil, e.NewInternalServerException(wrap)
	}

	return &records, nil
}

func (r *repositoryImpl) CreateAuthIdentity(uid string, email string, encryptedPassword string) (*AuthIdentity, e.Exception) {
	tx := r.db.Begin()

	user := User{Email: email}
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		wrap := errors.Wrap(err, "Failed to create User")
		return nil, e.NewInternalServerException(wrap)
	}

	authIdentity := &AuthIdentity{
		Provider:          ProviderPassword,
		UID:               uid,
		EncryptedPassword: encryptedPassword,

		User: user,
	}

	err = tx.Create(authIdentity).Error
	if err != nil {
		tx.Rollback()
		wrap := errors.Wrap(err, "Failed to create AuthIdentity")
		return nil, e.NewInternalServerException(wrap)
	}

	tx.Commit()
	return authIdentity, nil
}

func (r *repositoryImpl) FindAuthIdentityByUID(uid string) (*AuthIdentity, e.Exception) {
	aid := AuthIdentity{}
	err := r.db.Where("uid = ?", uid).First(&aid).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find AuthIdentity with UID")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewUnauthorizedException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	return &aid, nil
}
