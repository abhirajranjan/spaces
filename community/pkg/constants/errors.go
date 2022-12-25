package constants

const (
	BadRequestErr     = 400
	Forbidden         = 403
	Ok                = 200
	InternalServerErr = 500
	NoDataFound       = 404
	DataAlreadyExists = 409
)

var (
	Status_ErrDb                = NewStatus(InternalServerErr, "database error")
	Status_Ok                   = NewStatus(Ok, "")
	Status_NoDocuments          = NewStatus(NoDataFound, "no documents matched")
	Status_AccountAlreadyExists = NewStatus(DataAlreadyExists, "data already exists")
)

func GenerateBadRequest(reason string) *Status {
	return NewStatus(BadRequestErr, reason)
}
