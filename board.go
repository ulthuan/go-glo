package glo

import "time"

type Board struct {
	Id              string        `json:"id"`
	Name            string        `json:"name"`
	Columns         []BoardColumn `json:"columns"`
	ArchivedColumns []BoardColumn `json:"archived_columns"`
	InvitedMembers  []Member      `json:"invited_members"`
	Members         []Member      `json:"members"`
	ArchivedDate    time.Time     `json:"archived_date"`
	Labels          []Label       `json:"labels"`
	CreatedDate     time.Time     `json:"created_date"`
	CreatedBy       PartialUser   `json:"created_by"`
}
