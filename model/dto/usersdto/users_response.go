package usersdto

type UsersResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}
