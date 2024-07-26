package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(to, subject, body string) error {
	message := []byte("" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", os.Getenv("SMTPUSERNAME"), os.Getenv("SMTPPASSWORD"), os.Getenv("SMTPHOST"))

	err := smtp.SendMail(os.Getenv("SMTPHOST")+":"+os.Getenv("SMTPPORT"), auth, os.Getenv("SMTPUSERNAME"), []string{to}, message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
