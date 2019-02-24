package glo

import "time"

// BoardCard struct for response unmarshall
type BoardCard struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	Description        CardDescription `json:"description"`
	BoardID            string          `json:"board_id"`
	ColumnID           string          `json:"column_id"`
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

// CardDescription struct for response unmarshall
type CardDescription struct {
	Text        string      `json:"text"`
	CreatedDate time.Time   `json:"created_date"`
	UpdatedDate time.Time   `json:"updated_date"`
	CreatedBy   PartialUser `json:"created_by"`
	UpdatedBy   PartialUser `json:"updated_by"`
}

// CardDescriptionParam struct for request
type CardDescriptionParam struct {
	Text string `json:"text"`
}

// BoardCardParams struct for request
type BoardCardParams struct {
	Name        string               `json:"name"`
	Position    int                  `json:"position"`
	Description CardDescriptionParam `json:"description"`
	ColumnID    string               `json:"column_id"`
	Assignees   []PartialUser        `json:"assignees"`
	Labels      []PartialLabel       `json:"labels"`
	DueDate     time.Time            `json:"due_date"`
}
