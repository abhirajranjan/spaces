package constants

const (
	BadRequestErr     = 400
	Forbidden         = 403
	Ok                = 200
	InternalServerErr = 500
)

var (
	Status_ErrCql       = NewStatus(InternalServerErr, "database error")
	Status_ErrZeroQuery = NewStatus(InternalServerErr, "empty query passed")
	Status_Ok           = NewStatus(Ok, "")
)

func GenerateBadRequest(reason string) *Status {
	return NewStatus(BadRequestErr, reason)
}
