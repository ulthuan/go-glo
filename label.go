package glo

import (
	"image/color"
	"time"
)

// Label struct for response unmarshall
type Label struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Color       color.RGBA  `json:"color"`
	CreatedDate time.Time   `json:"created_date"`
	CreatedBy   PartialUser `json:"created_by"`
}

// PartialLabel struct for response unmarshall
type PartialLabel struct {
	ID string `json:"id"`
}
