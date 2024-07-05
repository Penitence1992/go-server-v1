package errors

import "fmt"

type ProxyError struct {
	HttpStatus int
	Message    string
}

func NewProxyErrors(msg string, status int) error {
	return &ProxyError{
		HttpStatus: status,
		Message:    msg,
	}
}

func (p *ProxyError) Error() string {
	return fmt.Sprintf("代理错误,状态码:%d, 消息:%s", p.HttpStatus, p.Message)
}
