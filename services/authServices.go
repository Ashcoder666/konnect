package services

import (
	"fmt"

	"github.com/ashcoder666/konnect/models"
)

// create user

func CreateUserService(user *models.User) error {

	if err := models.DB.Create(&user).Error; err != nil {
		if models.IsDuplicateKeyError(err) {
			return fmt.Errorf(" username or Email already exists")
		} else {
			return fmt.Errorf("error creating use")
		}
	} else {
		return nil
	}
}

// login user

// forgot password

// reset password
