package helpers

import (
	"encoding/json"
	"io"
	"strings"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

var blank interface{}

func (response *Response) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(response)
}

func SetAndGetResponse(success bool, message string, data interface{}, code int) string {
	var response interface{}
	resp := Response{Success: success, Message: message, Data: data, Code: code}

	successResponse, err := json.Marshal(resp)
	fatalResponse := ErrorResponse(err)

	if response = successResponse; len(fatalResponse) > 0 {
		response = fatalResponse
	}

	return strings.TrimSpace(string(response.([]byte)))
}

func ErrorResponse(err error) string {
	if err != nil {
		fatalResponse := SetAndGetResponse(false, err.Error(), blank, 400)
		return fatalResponse
	}
	return ""
}
