package views

type EventInfo struct {
	Id              int64  `json:"id"`
	Owner           string `json:"owner"`
	PartCount       int64  `json:"part_count"`
	Description     string `json:"description"`
}

type EventsInfo struct {
	Events          []EventInfo `json:"events"`
}
