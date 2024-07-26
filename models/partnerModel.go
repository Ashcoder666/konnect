package models

import (
	"time"

	"github.com/google/uuid"
)

type Partner struct {
	ID          uuid.UUID        `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()" json:"id"`
	UserName    string           `gorm:"unique; not null" json:"username"`
	Email       string           `gorm:"unique; not null" json:"email"`
	Password    string           `gorm:"not null" json:"password"`
	Role        UserRole         `gorm:"type:user_role;not null" json:"role"`
	CreatedAt   time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	BankDetails BankDetailsModel `gorm:"not null" json:"bank_details"`
}

type BankDetailsModel struct {
	BankAccountNumber int    `json:"bank_account_number"`
	BankName          string `json:"bank_name"`
	BankIFSC          string `json:"ifsc"`
}
