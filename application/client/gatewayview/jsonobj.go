package gatewayview

type EventInfo struct {
	Id              int64  `json:"id"`
	Owner           string `json:"owner"`
	PartCount       int64  `json:"part_count"`
	Description     string `json:"description"`
}

type EventsInfo struct {
	Events          []EventInfo `json:"events"`
}


type UserInfo struct {
	UserName string `json:"username"`
	CountEvens int64  `json:"count"`
	Id         int64  `json:"id"`
}

type RegistrationInfo struct {
	Id int64      `json:"id"`
	UserId string `json:"user_id"`
	EventId int64 `json:"event_id"`
}

type ClientInfo struct {
	Id 		int64 `json:"id"`
	Name 	string `json:"name"`
	RedURL string `json:"red_url"`
}

type RedirectInfo struct {
	RedirectUrl string `json:"redirect_url"`
	CodeFlow    string `json:"code_flow"`
}