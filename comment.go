package glo

import "time"

// Comment struct for response unmarshall
type Comment struct {
	ID          string      `json:"id"`
	CardID      string      `json:"card_id"`
	BoardID     string      `json:"board_id"`
	CreatedDate time.Time   `json:"created_date"`
	UpdatedDate time.Time   `json:"updated_date"`
	CreatedBy   PartialUser `json:"created_by"`
	UpdatedBy   PartialUser `json:"updated_by"`
	Text        string      `json:"text"`
}

// ParamComment struct for request
type ParamComment struct {
	Text string `json:"text"`
}
