package Helpers

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//http请求
func HttpRequest(api string,json string,method string) (string, error) {
	jsonStr := []byte(json)
	req, err := http.NewRequest(method, api, bytes.NewBuffer(jsonStr))
	//使用json格式传参
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", ApiServerError
	}

	//关闭请求体
	defer func() { _ = resp.Body.Close()}()

	body, _ := ioutil.ReadAll(resp.Body)

	if !(resp.StatusCode == 200) {
		return "", ApiServerError
	}
	return string(body), nil
}
