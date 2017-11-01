package views

type UserInfo struct {
	UserName string `json:"username"`
	CountEvens int64  `json:"count"`
	Id         int64  `json:"id"`
}
