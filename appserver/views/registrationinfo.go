package views

type RegistrationInfo struct {
	Id int64      `json:"id"`
	UserId string  `json:"user_id"`
	EventId int64 `json:"event_id"`
}

type AllRegInfo struct {
	Id int64 		`json:"id"`
	Event EventInfo `json:"event"`
	User  UserInfo  `json:"user"`
}