package views

type UserInfo struct {
	UserName string `json:"username"`
	CountEvens int64  `json:"count"`
	Id         int64  `json:"id"`
}

type LogIn struct {
	Login  string `json:"login"`
	Pass   string `json:"pass"`
} 