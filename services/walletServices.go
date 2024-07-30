package services

import (
	"fmt"

	"github.com/ashcoder666/konnect/models"
	"github.com/google/uuid"
)

func CreateWallet(user_id uuid.UUID) error {
	// check if user have a wallet already
	var userInstance models.User
	fmt.Println(user_id)
	var newWallet = models.Wallet{
		UserId:       user_id,
		WalletAmount: 0,
		Status:       true,
	}

	if err := models.DB.Where("user_id = ?", user_id).First(&userInstance); err == nil {
		return fmt.Errorf("user already exists")
	}

	// fmt.Println(userInstance)

	// return nil

	if err := models.DB.Create(&newWallet); err != nil {
		return err.Error
	}

	// if not then create one
	return nil
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

// socket ik
// webrtc is something new , but im not familiar
// lets
