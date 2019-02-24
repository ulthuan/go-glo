package glo

import "time"

type BoardCard struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	Description        CardDescription `json:"description"`
	BoardId            string          `json:"board_id"`
	ColumnId           string          `json:"column_id"`
	CreatedDate        time.Time       `json:"created_date"`
	UpdatedDate        time.Time       `json:"updated_date"`
	ArchivedDate       time.Time       `json:"archived_date"`
	Assignees          []PartialUser   `json:"assignees"`
	Labels             []PartialLabel  `json:"labels"`
	DueDate            time.Time       `json:"due_date"`
	CommentCount       int             `json:"comment_count"`
	AttachmentCount    int             `json:"attachment_count"`
	CompletedTaskCount int             `json:"completed_task_count"`
	TotalTaskCount     int             `json:"total_task_count"`
	CreatedBy          PartialUser     `json:"created_by"`
}

type CardDescription struct {
	Text        string      `json:"text"`
	CreatedDate time.Time   `json:"created_date"`
	UpdatedDate time.Time   `json:"updated_date"`
	CreatedBy   PartialUser `json:"created_by"`
	UpdatedBy   PartialUser `json:"updated_by"`
}

type CardDescriptionParam struct {
	Text string `json:"text"`
}

type BoardCardParams struct {
	Name        string               `json:"name"`
	Position    int                  `json:"position"`
	Description CardDescriptionParam `json:"description"`
	ColumnId    string               `json:"column_id"`
	Assignees   []PartialUser        `json:"assignees"`
	Labels      []PartialLabel       `json:"labels"`
	DueDate     time.Time            `json:"due_date"`
}
