package glo

import "time"

// Attachment struct for response unmarshall
type Attachment struct {
	ID          string    `json:"id"`
	Filename    string    `json:"filename"`
	MimeType    string    `json:"mimetype"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   Member    `json:"created_by"`
}
