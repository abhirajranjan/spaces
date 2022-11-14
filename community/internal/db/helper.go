package db

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
)

type Hex struct {
	value string
}
