package glo

import "time"

type Comment struct {
	Id          string      `json:"id"`
	CardId      string      `json:"card_id"`
	BoardId     string      `json:"board_id"`
	CreatedDate time.Time   `json:"created_date"`
	UpdatedDate time.Time   `json:"updated_date"`
	CreatedBy   PartialUser `json:"created_by"`
	UpdatedBy   PartialUser `json:"updated_by"`
	Text        string      `json:"text"`
}

type ParamComment struct {
	Text string `json:"text"`
}