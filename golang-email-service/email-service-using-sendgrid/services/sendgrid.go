package services

import (
	"fmt"
	"log"

	"../models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	// Key for SendGrid
	Key = "PUT_YOUR_SENDGRID_KEY_HERE"
)

// SendEmail function to send mail, will return true on success and false on failure
func SendEmail(recipients []models.Recipient, sender models.Sender, subject string, message string) bool {
	m := mail.NewV3Mail()

	from := mail.NewEmail(sender.SenderName, sender.SenderEmail)
	content := mail.NewContent("text/html", message)
	m.SetFrom(from)
	m.AddContent(content)

	personalization := mail.NewPersonalization()

	for _, recipient := range recipients {
		to := mail.NewEmail(recipient.RecipientName, recipient.RecipientEmail)
		switch recipient.RecipientCCBCC {
		case "CC":
			{
				personalization.AddCCs(to)
			}
		case "BCC":
			{
				personalization.AddBCCs(to)
			}
		default:
			{
				personalization.AddTos(to)
			}

		}
	}
	personalization.Subject = subject

	m.AddPersonalizations(personalization)
	request := sendgrid.GetRequest(Key, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return true
}
