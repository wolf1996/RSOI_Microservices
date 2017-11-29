package userclient

type UserDecrementMessage struct {
	UserId string `json:"user_id"`
}

type MessageTokened struct {
	Token string         `json:"token"`
	Message UserDecrementMessage `json:"message"`
}
