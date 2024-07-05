package errors

type CwError interface {
	error
	Code() int
	BizCode() string
	Data() interface{}
}

type BaseCwError struct {
	code    int
	data    interface{}
	bizCode string
	msg     string
}

func (b *BaseCwError) Code() int {
	return b.code
}

func (b *BaseCwError) Data() interface{} {
	return b.data
}

func (b *BaseCwError) Error() string {
	return b.msg
}

func (b *BaseCwError) BizCode() string {
	return b.bizCode
}

type ResourceNotFound struct {
	BaseCwError
}

func New(msg string) CwError {
	return &BaseCwError{
		msg:     msg,
		code:    500,
		bizCode: ServerError,
	}
}

func NewBaseCwError(code int, bizCode, msg string) CwError {
	return &BaseCwError{
		msg:     msg,
		code:    code,
		bizCode: bizCode,
	}
}

func NewServerError(msg string) *BaseCwError {
	return &BaseCwError{
		msg:     msg,
		code:    500,
		bizCode: ServerError,
	}
}

func NewResourceNotFoundError(msg string) *ResourceNotFound {
	return &ResourceNotFound{
		BaseCwError: BaseCwError{
			msg:     msg,
			code:    404,
			bizCode: RequestError,
		},
	}
}
