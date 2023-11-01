package opts

type Params struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type JsonResult struct {
	Code int         `json:"code"` // 0：success 1:faild
	Data interface{} `json:"data"` // 返回的data
}
