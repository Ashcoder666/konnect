package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ashcoder666/konnect/models"
	"github.com/ashcoder666/konnect/utils"
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

func SendEmail(to, subject, body string) error {
	if err := utils.SendMail(to, subject, body); err != nil {
		return err
	}

	return nil
}

// forgot password

func ForgotPassword(email string) int {
	// return otp and save in db

	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 99
	randomInt := rand.Intn(9999)
	newOtpInstance := models.Tempotps{
		Email: email, OTP: randomInt, Status: true,
	}

	if err := models.DB.Create(&newOtpInstance); err != nil {
		fmt.Println(err)
	}

	return randomInt
}

// reset password

func ResetPassword(email, newPassword string, OTP int) error {
	var otpInstance models.Tempotps
	var user models.User

	// get otp & check if its correcr

	err := models.DB.Where("email = ? AND status = true AND otp = ?", email, OTP).First(&otpInstance).Error

	if err != nil {
		return err
	}

	// delete record
	res := models.DB.Where("email = ? AND status = true AND otp = ?", email, OTP).Delete(&otpInstance)

	if res.Error != nil {
		fmt.Println("hryyyyyyyyyyyyyyyy")
		return res.Error
	}
	// update new password

	fmt.Println("line no 96")

	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("line no 105")
	fmt.Println(user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	if err := models.DB.Save(&user); err != nil {
		return err.Error
	}

	fmt.Println("line no 116")

	return nil

}

// partner registartion

func RegisterPartnerService(partner *models.Partner) error {
	if err := models.DB.Create(&partner).Error; err != nil {
		return err
	}

	return nil
}

// partner login

func PartnerExistCheckService(email string) (models.Partner, error) {
	var partnerInstance models.Partner

	if err := models.DB.Where("email = ?", email).First(&partnerInstance); err != nil {
		return models.Partner{}, err.Error
	}

	return partnerInstance, nil
}
