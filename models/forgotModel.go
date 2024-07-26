package models

type Tempotps struct {
	Email  string `json:"email"`
	OTP    int    `json:"otp"`
	Status bool   `json:"status"`
}
