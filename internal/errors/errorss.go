package errors

type TestError struct {
}

func (t *TestError) Error() string {
	return "test error"
}
