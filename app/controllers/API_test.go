package controllers

// import (
// 	"fmt"
// 	"net/http"
// )

// type MessageErr interface {
// 	Message() string
// 	Status() int
// 	Error() string
// }

// type messageErr struct {
// 	ErrMessage string `json:"message"`
// 	ErrStatus  int    `json:"status"`
// 	ErrError   string `json:"error"`
// }

// func (e *messageErr) Error() string {
// 	return e.ErrError
// 	fmt.Println()
// }

// func (e *messageErr) Message() string {
// 	return e.ErrMessage
// }

// func (e *messageErr) Status() int {
// 	return e.ErrStatus
// }

// func NewNotFoundError(message string) MessageErr {
// 	return &messageErr{
// 		ErrMessage: message,
// 		ErrStatus:  http.StatusNotFound,
// 		ErrError:   "not_found",
// 	}
// }
