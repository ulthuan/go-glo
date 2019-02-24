package glo

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PartialUser struct {
	Id string `json:"id"`
}
