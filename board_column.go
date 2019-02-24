package glo

import "time"

type BoardColumn struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	ArchivedDate time.Time `json:"archived_date"`
	CreatedDate  time.Time `json:"created_date"`
}

type CreateBoardColumnParams struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}
