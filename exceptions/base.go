package exceptions

type Error struct {
	errMsg string
}

func (e Error) Error() string {
	return e.errMsg
}

func NewErr(msg string) Error {
	return Error{msg}
}
