package glo

// User struct for response unmarshall
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// PartialUser struct for response unmarshall
type PartialUser struct {
	ID string `json:"id"`
}
