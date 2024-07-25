package services

import (
	"fmt"

	"github.com/ashcoder666/konnect/models"
	"golang.org/x/crypto/bcrypt"
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

func CheckUserExists(emailorusername string) (*models.User, error) {
	var user *models.User
	err := models.DB.Where("email = ? OR user_name = ?", emailorusername, emailorusername).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func ComparePassword(inputPassword, cryptPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(cryptPassword), []byte(inputPassword)); err != nil {
		return err
	}

	return nil
}

func SendEmail() {

}

// forgot password

func ForgotPassword(email string) {

}

// reset password
