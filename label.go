package glo

import (
	"image/color"
	"time"
)

type Label struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Color       color.RGBA  `json:"color"`
	CreatedDate time.Time   `json:"created_date"`
	CreatedBy   PartialUser `json:"created_by"`
}

type PartialLabel struct {
	Id string `json:"id"`
}