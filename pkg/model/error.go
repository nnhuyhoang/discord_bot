package model

type InternalError struct {
}

func (e *InternalError) Error() string {
	return "Internal Server Error"
}
