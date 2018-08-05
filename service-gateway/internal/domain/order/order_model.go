package order

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

type Order struct {
	persistent.BaseModel

	State  string `gorm:"column:state; 	type: VARCHAR(30); 		NOT NULL;"`
	Amount uint   `gorm:"column:amount; type:UNSIGNED BIG INT;	NOT NULL;"`

	ShippingCountry  string `gorm:"column:shipping_country; 	type:VARCHAR(50); 	NOT NULL;"`
	ShippingCity     string `gorm:"column:shipping_city; 		type:VARCHAR(50); 	NOT NULL;"`
	ShippingState    string `gorm:"column:shipping_state; 		type:VARCHAR(50); 	NOT NULL;"`
	ShippingZipCode  string `gorm:"column:shipping_zipcode; 	type:VARCHAR(20); 	NOT NULL;"`
	ShippingAddress1 string `gorm:"column:shipping_address1; 	type:TEXT; 			NOT NULL;"`
	ShippingAddress2 string `gorm:"column:shipping_address2; 	type:TEXT; 			NOT NULL;"`
	ShippingMessage  string `gorm:"column:shipping_message; 	type:TEXT; 			NOT NULL;"`

	OrdererName  string `gorm:"column:orderer_name; 	type:VARCHAR(50); 	NOT NULL;"`
	OrdererPhone string `gorm:"column:orderer_phone; 	type:VARCHAR(50); 	NOT NULL;"`
	OrdererEmail string `gorm:"column:orderer_email; 	type:VARCHAR(50); 	NOT NULL;"`

	RecipientName  string `gorm:"column:recipient_name; 	type:VARCHAR(50); 	NOT NULL;"`
	RecipientPhone string `gorm:"column:recipient_phone; 	type:VARCHAR(50); 	NOT NULL;"`
	RecipientEmail string `gorm:"column:recipient_email; 	type:VARCHAR(50); 	NOT NULL;"`

	UserID uint `gorm:"column:user_id" sql:"type:UNSIGNED BIG INT REFERENCES User(id) ON DELETE RESTRICT ON UPDATE CASCADE"`

	OrderDetails []Order `gorm:"foreignkey:OrderID"`
}

func (Order) TableName() string {
	return "Order"
}
