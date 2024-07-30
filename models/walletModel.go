package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID           uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	UserId       uuid.UUID `gorm:"unique;type:uuid" json:"user_id"`
	WalletAmount int       `gorm:"default:0" json:"wallet_amount"`
	Status       bool      `json:"status"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
}
