package glo

import "time"

// BoardColumn struct for response unmarshall
type BoardColumn struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ArchivedDate time.Time `json:"archived_date"`
	CreatedDate  time.Time `json:"created_date"`
}

// CreateBoardColumnParams struct for request
type CreateBoardColumnParams struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}
