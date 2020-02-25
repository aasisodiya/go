package models

// Recipient struct
type Recipient struct {
	RecipientName  string
	RecipientEmail string
	RecipientCCBCC string // Use CC / BCC as value
}

// Sender struct
type Sender struct {
	SenderName  string
	SenderEmail string
}
