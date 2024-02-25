package model

type MsgInfo struct {
	ToUser  string `p:"touser" `
	AgentId string `p:"agentid" `
	Content string `p:"content" `
}
type AllMsg struct {
	ToUser  string   `json:"touser" `
	Toparty string   `json:"toparty" `
	Totag   string   `json:"totag" `
	Msgtype string   `json:"msgtype" `
	Text    *MsgType `json:"text" `
	AgentId string   `json:"agentid" `
}
type MsgType struct {
	Content string `json:"content" `
}
