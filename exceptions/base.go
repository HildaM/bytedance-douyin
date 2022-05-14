package exceptions

type Error struct {
	ErrMsg string
}

func (e Error) Error() string {
	return e.ErrMsg
}

func NewErr(msg string) Error {
	return Error{msg}
}
