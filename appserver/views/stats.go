package views

type LoginEvent struct {
	Ok 	 	 bool	   `json:"ok" bson:"ok"`
	Info     string    `json:"info" bson:"info"`
}

type ViewEvent struct {
	Path 	 string	   `json:"path" bson:"path"`
	UserId   string	   `json:"user_id" bson:"user_id"`
}

type ChangeEvent struct {
	Path 	 string	   `json:"path" bson:"path"`
	UserId   string	   `json:"user_id" bson:"user_id"`
}
