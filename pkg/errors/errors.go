package errors

import(
	"github.com/daniel-oc/podcast-clipper/pkg/models"
)

func NewCustomError(status int, name string, message string) models.CustomError {
	return models.CustomError{
		Status: status,
		Name: name,
		Message: message,
	}
}

func NewInvalidURLError(message string) models.CustomError {
	return NewCustomError(400, "InvalidURL", message)
}

func NewScrapingError(message string) models.CustomError {
    return NewCustomError(500, "ScrapingError", message)
}

func WrapError(err error, status int, name string) models.CustomError {
    return models.CustomError{
        Status:  status,
        Name:    name,
        Message: err.Error(),
    }
}