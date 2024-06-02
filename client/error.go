package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiwasaki126/chat_gpt_api_usage_exporter/client/response"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (e *Error) Error() string {
	return e.Message
}

func newBadRequestError(resp response.ErrorResponse) *Error {
	return &Error{
		Code:    http.StatusBadRequest,
		Message: *resp.Error.Message,
		Type:    *resp.Error.Type,
	}
}

func newInternalServerErrorInClient(err error) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		Type:    "internal_server_error",
	}
}

func newInternalServerErrorInAPI(resp response.ErrorResponse) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: *resp.Error.Message,
		Type:    *resp.Error.Type,
	}
}

func handleErrorResponse(resp *http.Response) error {
	var body response.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return newInternalServerErrorInClient(err)
	}
	if err := body.Validate(nil); err != nil {
		return newInternalServerErrorInClient(err)
	}
	switch resp.StatusCode {
	case http.StatusBadRequest:
		return newBadRequestError(body)
	case http.StatusInternalServerError:
		return newInternalServerErrorInAPI(body)
	default:
		return newInternalServerErrorInClient(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}
}
