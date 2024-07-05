package errors

var (
	ErrResourceNotFound = NewResourceNotFoundError("资源不存在")
	ErrServerException  = NewServerError("服务器异常")
)

const (
	RequestError = "A0001"
	ServerError  = "B0001"
)
