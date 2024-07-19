package models

type CustomError struct {
	Status int `json:"status"`
	Name string `json:"name"`
	Message string `json:"message"`
}

func (e CustomError) Error() string {
	return e.Message
}