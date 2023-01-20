package smtps

import (
	"fmt"
	"net/smtp"
)

func Smtp(to string, message string) {
	from := "aryanpal692@gmail.com"
	password := "uhmuvyuxufshuurw"

	smtpHost := "smtp.gmail.com"
	smtpPort := "25"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}
