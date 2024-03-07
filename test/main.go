package main

import (
	"fmt"
	"log"

	"github.com/linxGnu/gosmpp"
	"github.com/linxGnu/gosmpp/data"
	"github.com/linxGnu/gosmpp/pdu"
)

func main() {
	// Define the SMPP client options
	opts := gosmpp.SessionOpts{
		Host:           "smsc.example.com",
		Port:           2775,
		SystemID:       "username",
		Password:       "password",
		SystemType:     "",
		EnquireLink:     60,
		SubmitTimeout:  10 * 1000,
		SubmitRetries:  3,
		BindInterval:   5 * 1000,
		Handler:        pdu.HandlerFunc(handlePDU),
		ReconnectDelay: 5,
	}
	session, err := gosmpp.NewSession(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Define the message payload and destination number
	message := "Hello, World!"
	destination := "1234567890"

	// Create a new SubmitSM request with the message payload and destination number
	sm := pdu.NewSubmitSM()
	sm.SourceAddr = "SMPPTest"
	sm.DestinationAddr = destination
	sm.ShortMessage = []byte(message)
	sm.RegisteredDelivery = data.SMRegisteredDeliveryReceipt

	// Send the SubmitSM request to the SMSC and wait for the response
	resp, err := session.Submit(sm)
	if err != nil {
		log.Fatal(err)
	}

	// Check the response status and print the message ID if successful
	if resp.Header().ID == data.SUBMIT_SM_RESP && resp.Status() == data.ESME_ROK {
		msgID := resp.GetField(data.FiledMessageID)
		fmt.Println("Message sent successfully with ID:", msgID.Value())
	} else {
		log.Fatalf("Failed to send message. Status: %s", resp.Status())
	}
}

// Define a function to handle incoming PDUs (optional)
func handlePDU(p pdu.PDU, ses *gosmpp.Session) error {
	log.Println("Received PDU:", p.Header().ID)
	return nil
}
