package views

type RegistrationInfo struct {
	Id int64
	UserId int64
	EventId int64
}

type AllRegInfo struct {
	Id int64
	Event EventInfo
	User  UserInfo
}