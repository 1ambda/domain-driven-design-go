package domain

import (
	"github.com/jinzhu/gorm"

	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/satori/go.uuid"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/config"
)

func Transact(db *gorm.DB, callback func(tx *gorm.DB) (e.Exception)) (e.Exception) {
	u := uuid.NewV4().String()
	logger := config.GetDbLogger()

	tx := db.Begin()
	if tx.Error != nil {
		ex := e.NewInternalServerException(tx.Error, "Unknown transaction error: Can't start")
		return ex
	}

	if config.Env.DebugTransactionEnabled() {
		logger.Infow("Transaction started", "uuid", u)
	}

	ex := callback(tx)
	if ex != nil {
		tx.Rollback()
		if config.Env.DebugTransactionEnabled() {
			logger.Infow("Transaction rollback-ed", "uuid", u)
		}
		return ex
	}

	tx.Commit()
	if tx.Error != nil {
		ex := e.NewInternalServerException(tx.Error, "Unknown transaction error: Can't commit")
		return ex
	}

	if config.Env.DebugTransactionEnabled() {
		logger.Infow("Transaction committed", "uuid", u)
	}

	return nil
}
