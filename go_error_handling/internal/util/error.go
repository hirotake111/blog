package util

type AppError struct {
	Status int
	detail string
}

func (e AppError) Error() string {
	return e.detail
}

func (e AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return t.Status == e.Status && t.detail == e.detail
}

var BadRequestError = AppError{Status: 400, detail: "bad Request"}
var InternalServerError = AppError{Status: 500, detail: "internal server request"}
