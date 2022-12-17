package db

import "log"

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	AccountAlreadyExists = Error("AccountAlreadyExists")
	InvalidPrams         = Error("InvalidPrams")
	NoAccountExists      = Error("NoAccountExists")
	NameCannotBeNull     = Error("NameCannotBeNull")
	DescCannotBeNull     = Error("DescCannotBeNull")
	TagCannotBeNull      = Error("TagCannotBeNull")
	DBError              = Error("DBError")
)

type Hex struct {
	value string
}

func handleError(prefix string, err error) {
	log.SetPrefix(prefix)
	log.Panic(err)
}
