package ikuai

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Response[T any] struct {
	Result int    `json:"Result"`
	ErrMsg string `json:"ErrMsg"`
	Data   *T     `json:"Data"` // This can hold any type of data for dynamic responses
	RowId  *int   `json:"RowId"`
}

const (
	FuncNameParam     = "func_name"
	ActionParam       = "action"
	ParamParam        = "param"
	ActionCallPath    = "/Action/call"
	HeaderCookie      = "Cookie"
	HeaderContentType = "Content-Type"
)

func actionFailed[T any](response *Response[T]) bool {
	return response.Result == 10014 || response.ErrMsg == "no login authentication"
}

func CallAction[T any, P any](funcName string, action string, param P) (r *Response[T], err error) {
	var jsonBytes []byte
	var request *http.Request
	var response *http.Response
	var result = Response[T]{}
	var router *RpcClient
	// Create a new response object
	router, err = DefaultClient()
	if err != nil {
		return
	}
	actionUrl := router.config.Url + ActionCallPath
	params := map[string]any{
		FuncNameParam: funcName,
		ActionParam:   action,
		ParamParam:    param,
	}
	// Marshal the params to json
	if jsonBytes, err = json.Marshal(params); err != nil {
		return
	}
	if request, err = http.NewRequest(http.MethodPost, actionUrl, bytes.NewBuffer(jsonBytes)); err != nil {
		return
	}
	// Set the request headers
	header := request.Header
	header.Set(HeaderCookie, router.cookie())
	header.Set(HeaderContentType, ContentTypeApplicationJSON)
	// Send the request
	if response, err = router.client.Do(request); err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(response.Body)
	// If the response status code is 502 Bad Gateway, try to log in again
	if response.StatusCode == http.StatusBadGateway {
		if err = router.login(); err != nil {
			return
		}
		return CallAction[T, P](funcName, action, param)
	}
	// Read the response body
	if jsonBytes, err = io.ReadAll(response.Body); err != nil {
		return
	}
	if err = json.Unmarshal(jsonBytes, &result); err != nil {
		return
	}
	// If the action failed and retry is enabled, try to log in again
	if router.config.Retry.Enable && actionFailed(&result) {
		if err = router.login(); err != nil {
			return
		}
		return CallAction[T, P](funcName, action, param)
	}
	// Log the action call
	if router.config.Log {
		router.Printf("Called action [%s] param=%v -> response: %s\n", action, param, jsonBytes)
	}
	return &result, err
}
