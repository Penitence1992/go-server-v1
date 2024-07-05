package api

type CwResponse struct {
	BizCode string      `json:"bizCode"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

func Ok(data interface{}) *CwResponse {
	return &CwResponse{
		BizCode: "",
		Code:    200,
		Data:    data,
		Msg:     "",
	}
}

func Error(code int, bizCode, msg string) *CwResponse {
	return &CwResponse{
		BizCode: bizCode,
		Code:    code,
		Data:    nil,
		Msg:     msg,
	}
}
