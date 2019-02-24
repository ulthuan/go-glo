package glo

// Member struct for response unmarshall
type Member struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
}
