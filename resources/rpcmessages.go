package resources

type MessageTokened struct {
	Token string         `json:"token"`
	Message interface{}  `json:"message"`
}
