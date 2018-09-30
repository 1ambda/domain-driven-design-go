package user

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

const ProviderPassword = "password"
const ProvideGithub = "github"

type AuthIdentity struct {
	persistent.BaseModel

	Provider          string `gorm:"column:provider; type:VARCHAR(255); NOT NULL;"`
	UID               string `gorm:"column:uid; type:VARCHAR(20); UNIQUE; INDEX; NOT NULL;"`
	EncryptedPassword string `gorm:"column:encrypted_password; type:TEXT; NOT NULL;"`

	User   User `gorm:"foreignkey:UserID;"`
	UserID uint `gorm:"olumn:user_id" sql:"type:UNSIGNED BIG INT REFERENCES User(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
}

func (AuthIdentity) TableName() string {
	return "AuthIdentity"
}

type AuthClaim struct {
	Provider string
	UserID   uint
	UID      string
}

func (aid *AuthIdentity) ToClaims() *AuthClaim {
	return &AuthClaim{
		Provider: aid.Provider,
		UserID:   aid.UserID,
		UID:      aid.UID,
	}
}
