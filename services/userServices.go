package services

import (
	"fmt"

	"github.com/ashcoder666/konnect/models"
)

func ListAllPartners() ([]models.Partner, error) {
	var partnerInstance []models.Partner

	if err := models.DB.Find(&partnerInstance).Error; err != nil {
		fmt.Println("Error fetching partners:", err)
		return nil, err
	}

	return partnerInstance, nil

}
