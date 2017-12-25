package views

type ClientInfo struct {
	Id 		int64 `json:"id"`
	Name 	string `json:"name"`
	RedURL string `json:"red_url"`
}

type RedirectInfo struct {
	RedirectUrl string `json:"redirect_url"`
	CodeFlow    string `json:"code_flow"`
}

type CodeFlowView struct {
	CodeFlow string `json:"code_flow"`
	Domain   string `json:"domain"`
}