package errors

import "strconv"

type Error struct {
	Code    int
	Name    string
	Message string
}

func IsEmptyError(err Error) bool {
	return err.Code == 0 && err.Name == "" && err.Message == ""
}

func (e Error) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Message
}

func (e Error) String() string {
	return e.Name + ": " + e.Message
}
