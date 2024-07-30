package services

import (
	"fmt"

	"github.com/ashcoder666/konnect/models"
	"github.com/google/uuid"
)

func CreateWallet(user_id uuid.UUID) error {
	// check if user have a wallet already
	var userInstance models.User

	if err := models.DB.Where("user_id = ?", user_id).First(&userInstance); err != nil {
		return err.Error
	}

	fmt.Println(userInstance)

	return nil

	// if not then create one
}

func AddMoney(Amount int, userid uuid.UUID) {
	// find wallet

	// update it

	// transaction entry
}

func DeductMoney(Amount int, userid uuid.UUID) {
	// find wallet

	// update it

	// transaction entry
}

func AddTransactionService(WalletId, UserId uuid.UUID, PaymentReferenceNumber, PaymentMethod string) {
	// add transcation entry
}
