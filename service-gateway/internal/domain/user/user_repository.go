package user

import (
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	DeleteUser(id uint) (bool, e.Exception)
	FindUserById(id uint) (*User, e.Exception)
	FindUserByIdWithTx(tx *gorm.DB, id uint) (*User, e.Exception)
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
		return false, e.NewInternalServerException(result.Error, "Failed to delete User")
	}

	if result.RowsAffected < 1 {
		err := errors.New("Failed to fine User to be deleted")
		return false, e.NewNotFoundException(err, "Failed to find User which does not exist")
	}

	return true, nil
}

func (r *repositoryImpl) FindUserByIdWithTx(tx *gorm.DB, id uint) (*User, e.Exception) {
	record := &User{}
	err := tx.Where("id = ?", id).First(record).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(err, "Failed to find User does not exist")
		}

		return nil, e.NewInternalServerException(err, "Failed to find User")
	}

	return record, nil
}

func (r *repositoryImpl) FindUserById(id uint) (*User, e.Exception) {
	tx := r.db.Begin()
	if tx.Error != nil {
		ex := e.NewInternalServerException(tx.Error, "Failed to get transaction")
		return nil, ex
	}

	record, ex := r.FindUserByIdWithTx(tx, id)
	if ex != nil {
		return nil, ex
	}

	tx.Commit()
	if tx.Error != nil {
		ex := e.NewInternalServerException(tx.Error, "Failed to commit transaction")
		return nil, ex
	}

	return record, nil
}

func (r *repositoryImpl) FineAllUsers() (*[]User, e.Exception) {
	// TODO: use db.tx

	var records []User

	err := r.db.Find(&records).Error
	if err != nil {
		return nil, e.NewInternalServerException(err, "Failed to find all User")
	}

	return &records, nil
}

func (r *repositoryImpl) CreateAuthIdentity(uid string, email string, encryptedPassword string) (*AuthIdentity, e.Exception) {
	tx := r.db.Begin()

	user := User{Email: email}
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, e.NewInternalServerException(err, "Failed to create User")
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
		return nil, e.NewInternalServerException(err, "Failed to create AuthIdentity")
	}

	tx.Commit()
	return authIdentity, nil
}

func (r *repositoryImpl) FindAuthIdentityByUID(uid string) (*AuthIdentity, e.Exception) {
	aid := AuthIdentity{}
	err := r.db.Where("uid = ?", uid).First(&aid).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewBadRequestException(err, "Failed to find AuthIdentity doest not exist")
		}

		return nil, e.NewInternalServerException(err, "Failed to find AuthIdentity")
	}

	user := User{}
	aid.User = user

	err = r.db.Model(&aid).Related(&aid.User).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewBadRequestException(err, "Failed to find User by AuthIdentity which does not exist")
		}

		return nil, e.NewInternalServerException(err, "Failed to find User by AuthIdentity")
	}

	return &aid, nil
}
