package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                     uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	WalletId               uuid.UUID `gorm:"type:uuid" json:"wallet_id"`
	PaymentReferenceNumber string    `gorm:"unique; not null" json:"payment_reference_number"`
	PaymentMethod          string    `gorm:"unique; not null" json:"payment_method"`
	UserId                 uuid.UUID `gorm:"type:uuid" json:"user_id"`
	CreatedAt              time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
