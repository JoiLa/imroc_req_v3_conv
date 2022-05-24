package imroc_req_v3_conv

import (
	"net/http"

	v3 "github.com/imroc/req/v3"
)

// 将请求响应转`[]byte`
func ReqResponseConvToBytes(receiveResponse *v3.Response, receiveErr error) ([]byte, error) {
	if receiveErr != nil {
		return []byte{}, receiveErr
	}
	return receiveResponse.Bytes(), nil
}

// 将请求响应转`[]byte`，并且获取响应头
func ReqResponseConvToBytesAndGetHeaders(receiveResponse *v3.Response, receiveErr error) (http.Header, []byte, error) {
	if receiveErr != nil {
		return nil, []byte{}, receiveErr
	}
	return receiveResponse.Header, receiveResponse.Bytes(), nil
}

// 将请求响应转`string`
func ReqResponseConvToString(receiveResponse *v3.Response, receiveErr error) (string, error) {
	if receiveErr != nil {
		return "", receiveErr
	}
	return receiveResponse.String(), nil
}
