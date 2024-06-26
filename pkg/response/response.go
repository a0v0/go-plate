package response

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Alias for any slice.
type Messages = []any

// A struct to handle error with custom error handler.
type Error struct {
	Code     int   `json:"code"`
	Messages []any `json:"messages"`
}

// Error makes it compatible with the `error` interface.
func (e *Error) Error() string {
	return fmt.Sprint(e.Messages)
}

// A struct to return normal responses.
type Response struct {
	Code     int      `json:"code"`
	Messages Messages `json:"messages,omitempty"`
	Data     any      `json:"data,omitempty"`
}

// Nothing to describe this fucking variable.
var IsProduction bool

// Default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Response{
		Code: fiber.StatusInternalServerError,
	}
	// Handle errors
	if e, ok := err.(validator.ValidationErrors); ok {
		resp.Code = fiber.StatusForbidden
		resp.Messages = Messages{removeTopStruct(e.Translate(trans))}
	} else if e, ok := err.(*fiber.Error); ok {
		resp.Code = e.Code
		resp.Messages = Messages{e.Message}
	} else if e, ok := err.(*Error); ok {
		resp.Code = e.Code
		resp.Messages = Messages{e.Messages}

		// for ent and some errors
		if resp.Messages == nil {
			resp.Messages = Messages{err}
		}
		return Resp(c, resp)
	} else {
		resp.Messages = Messages{err.Error()}
	}

	if !IsProduction {
		log.Printf("From: Fiber's error handler. \n Error: %v", err)
	}

	return Resp(c, resp)
}

// NewErrors creates multiple new Error messages
func NewErrors(code int, messages ...any) *Error {
	e := &Error{
		Code:     code,
		Messages: []any{utils.StatusMessage(code)},
	}
	if len(messages) > 0 {
		e.Messages = messages
	}
	return e
}

// NewError creates singular new Error message
func NewError(code int, messages ...any) *Error {
	e := &Error{
		Code: code,
	}
	if len(messages) > 0 {
		e.Messages = messages
	} else {
		e.Messages = []any{utils.StatusMessage(code)}
	}
	return e
}

// A fuction to return beautiful responses.
func Resp(c *fiber.Ctx, resp Response) error {
	// Set status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}

	c.Status(resp.Code)

	// Return JSON
	return c.JSON(resp)
}

// Remove unnecessary fields from validator message
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, msg := range fields {
		stripStruct := field[strings.Index(field, ".")+1:]
		//res[stripStruct] = strings.TrimLeft(msg, stripStruct)
		res[stripStruct] = msg
	}
	return res
}
