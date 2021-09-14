package api

type RespBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RequestAgentState struct {
	AppId   string `json:"appId"`
	AgentId string `json:"agentId"`
}

type ResponseAgentState struct {
	RespBase
	Data: struct {
		AppId   string `json:"appId"`
		AgentId string `json:"agentId"`
	} `json:"data"`
}
