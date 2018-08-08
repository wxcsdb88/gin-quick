package common

import (
	"fmt"
)

// ResponseData response data
type ResponseData struct {
	Code int         `json:"code"` // http status code
	Msg  string      `json:"msg"`  // error msg
	Data interface{} `json:"data"` // data, can include data code
	URI  string      `json:"uri"`  // request uri
}

// NewResponseData get the ResponseData according the fields
func NewResponseData(code int, msgData interface{}, data interface{}, uri string) *ResponseData {
	var msg string
	if msgData != nil {
		switch mg := msgData.(type) {
		case error:
			msg = mg.Error()
		case string:
			msg = mg
		default:
			msg = fmt.Sprintf("%v", mg)
		}
	}

	return &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
		URI:  uri,
	}
}
