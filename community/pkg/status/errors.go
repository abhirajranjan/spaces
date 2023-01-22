package status

import "fmt"

const (
	BadRequestErrCode     = 400
	Forbidden             = 403
	OkCode                = 200
	InternalServerErrCode = 500
	NoDataFoundCode       = 404
	DataAlreadyExistsCode = 409
)

type Status struct {
	Value int
	Err   *error
}

func (d *Status) Error() string {
	return fmt.Sprint(d.Err)
}

var (
	ErrDb             = NewStatus(InternalServerErrCode, "database error")
	Ok                = NewStatus(OkCode, "")
	NoDataFound       = NewStatus(NoDataFoundCode, "no documents matched")
	DataAlreadyExists = NewStatus(DataAlreadyExistsCode, "data already exists")
)

func GenerateBadRequest(reason string) *Status {
	return NewStatus(BadRequestErrCode, reason)
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
