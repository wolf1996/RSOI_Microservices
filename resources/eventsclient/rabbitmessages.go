package eventsclient

type DecrementRegistration struct {
	EventId int64 `json:"event_id"`
}

type MessageTokened struct {
	Token string         `json:"token"`
	Message DecrementRegistration  `json:"message"`
}
