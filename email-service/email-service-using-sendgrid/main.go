package main

// Step 1: Import Models and S
import (
	"./models"
	"./services"
)

// Example demonstrating how to use the SendEmail Function
func main() {
	// ! Important: make sure you have provided correct key inside sendgrid.go constant

	// Step 1 : Create one Sender Object
	var s models.Sender
	s.SenderEmail = "example@test.com"
	s.SenderName = "SendGrid"

	// Step 2 : Create Required number of Recipients Objects
	var r1 models.Recipient
	r1.RecipientName = "Akash"
	r1.RecipientEmail = "akash_sisodiya@test.com"
	r1.RecipientCCBCC = "none"

	// // Just an example... uncomment if required
	// var r2 models.Recipient
	// r2.RecipientName = "Tiger"
	// r2.RecipientEmail = "aditya_sisodiya@test.com"
	// r2.RecipientCCBCC = "CC"

	// Step 3 : Create an Array of Recipient Object and populate the array
	var r []models.Recipient
	r = append(r, r1)
	// r = append(r, r2) // Uncomment this if second recipient is in play

	// Step 4 : Your are ready to send an email, just modify the data below as per your requirements
	subject := "Sample Subject"
	plainTextMessage := "This is Test Message!"
	htmlTextMessage := "<h1>This is Test Message!</h1>"

	// Sending one sample Email
	services.SendEmail(r, s, subject, plainTextMessage)
}
