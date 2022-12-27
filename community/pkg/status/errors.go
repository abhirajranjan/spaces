package status

import "fmt"

const (
	BadRequestErr     = 400
	Forbidden         = 403
	Ok                = 200
	InternalServerErr = 500
	NoDataFound       = 404
	DataAlreadyExists = 409
)

type Status struct {
	Value int
	Err   *error
}

func (d *Status) Error() string {
	return fmt.Sprint(d.Err)
}

var (
	ErrDb                = NewStatus(InternalServerErr, "database error")
	OkStatus             = NewStatus(Ok, "")
	NoDocuments          = NewStatus(NoDataFound, "no documents matched")
	AccountAlreadyExists = NewStatus(DataAlreadyExists, "data already exists")
)

func GenerateBadRequest(reason string) *Status {
	return NewStatus(BadRequestErr, reason)
}

func NewStatus(value int, v any) *Status {
	var err error
	switch cus := v.(type) {
	case string:
		err = fmt.Errorf(cus)
	case error:
		err = cus
	default:
		return nil
	}
	return &Status{
		Value: value,
		Err:   &err,
	}
}
