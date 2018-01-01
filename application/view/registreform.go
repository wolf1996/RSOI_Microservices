package view

type Registre struct {
	EventId int64 `form:"event" json:"event" binding:"required"`
}

type RemoveReg struct {
	RegId int64 `form:"reg" json:"reg" binding:"required"`
}

type LoginForm struct {
	Login    string `form:"login" binding:"required"`
	Password string `form:"password" binding:"required"`
}