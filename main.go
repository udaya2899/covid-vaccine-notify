package main

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/udaya2899/covid-vaccine-notify/model"
	"github.com/udaya2899/covid-vaccine-notify/services"
)

func main() {
	doEvery(time.Minute, getCenters)
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func getCenters(time.Time) {
	res, err := services.GetCenters("571", "01-05-2021")
	if err != nil {
		log.Fatal(err)
	}

	if len(res.Centers) == 0 {
		log.Printf("Centers not found for given district and date")
		return
	}

	err = sendEmail(res)
	if err != nil {
		log.Fatal(err)
	}
}

func sendEmail(res model.CowinResponse) error {
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		from,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	j, err := json.Marshal(res)
	if err != nil {
		return err
	}
	// Message.
	message := j

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	log.Printf("Email Sent Successfully!")
	return nil
}
