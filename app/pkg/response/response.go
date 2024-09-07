package response

import (
	"encoding/json"
	"net/http"
)

type CustomResponse struct {
	Headers Headers     `json:"headers"`
	Data    interface{} `json:"data"`
}

type Headers struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var (
	SuccessHeader = Headers{
		Status:  "success",
		Message: "Request was completed successfully",
		Code:    200,
	}
)

func ErrorHeader(message string, code int) Headers {
	return Headers{
		Status:  "error",
		Message: message,
		Code:    code,
	}
}

func RespondWithCustomResponse(w http.ResponseWriter, headers Headers, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(headers.Code)

	response := CustomResponse{
		Headers: headers,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}
