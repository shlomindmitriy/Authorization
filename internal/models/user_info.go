package models

type UserInfo struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	//LogName  string `json:"log_name"`
	Password string `json:"password"`
}
