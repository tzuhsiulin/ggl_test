package customerror

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

func NewErr(errCode int, args ...string) *CustomError {
	customErr := &CustomError{
		ErrorCode: errCode,
	}
	if len(args) > 0 {
		customErr.ErrorMsg = args[0]
	}
	return customErr
}
