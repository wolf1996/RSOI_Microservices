package shared

var (
	ResponceExchangeName = "response"
	StatisticExchangeName = "statistic"
	TypeGetInf = "getinf"
	TypeChangeInf = "changeinf"
	TypeLogin = "logins"
)

type MessageId struct {
	Producer  string `json:"producer" bson:"producer"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
	Random    int  `json:"random" bson:"random"`
	MsgType   string `json:"msg_type" bson:"msg_type"`
}

type ResponceMsg struct {
	ResponceStatus int64 `json:"responce_status"`
	MessageIds     MessageId `json:"message_ids"`
}

type  InfoViewMsg struct {
	Id 		 MessageId `json:"message_ids" bson:"message_ids"`
	Path 	 string	   `json:"path" bson:"path"`
	UserId   int64	   `json:"user_id" bson:"user_id"`
}

type  InfoChangeMsg struct {
	Id 		 MessageId `json:"message_ids" bson:"message_ids"`
	Path 	 string	   `json:"path" bson:"path"`
	UserId   int64	   `json:"user_id" bson:"user_id"`
}

type  LoginMsg struct {
	Id 		 MessageId `json:"message_ids" bson:"message_ids"`
	Ok 	 	 bool	   `json:"ok" bson:"ok"`
	Info     string    `json:"info" bson:"info"`
}