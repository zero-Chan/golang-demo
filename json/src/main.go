package main

import (
	"encoding/json"
	"fmt"
	//	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ParameterError struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields"`
}

type GetBusinessError struct {
	Errors  []string `json:"errors"`
	Message string   `json:"message"`
}

func main() {
	vals := make(url.Values)
	vals.Add("dynamic", "true")
	vals.Add("id", "142034145494497")

	b, err := Get("http://mediamax.quantone.com/creatives/status", vals)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	//	body := []byte(`
	//		{"errors":["202171"],"message":"Validation Error"}
	//	`)

	//	err := checkGetAPIError(body)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

}

var Errors = map[string]string{
	"202171": "创意不存在",
}

func Get(_url string, query url.Values) (body []byte, err error) {
	//	urlStr := _url + "?" + query.Encode()
	urlStr := "http://mediamax.quantone.com/creatives/status?dynamic=true&id=142034145494497"

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Basic bWVkaWFtYXgtYmlkZGluZ3g6U3VuMjAxNHRlbmc=")
	fmt.Printf("req: %s, header:%+v\n", urlStr, req.Header)

	cli := http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("异常的http状态码: %d", resp.StatusCode)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	fmt.Println("resp ", string(body))

	err = checkGetAPIError(body)
	if err != nil {
		return
	}

	return body, nil
}

func checkGetAPIError(body []byte) error {
	perror := new(ParameterError)
	err := json.Unmarshal(body, perror)
	if err == nil && len(perror.Fields) > 0 {
		return fmt.Errorf("参数错误: %+v", perror)
	}

	berror := new(GetBusinessError)
	err = json.Unmarshal(body, berror)
	if err == nil && len(berror.Errors) > 0 {
		var errmsgs []string
		for _, errcode := range berror.Errors {
			if msg, ok := Errors[errcode]; ok {
				errmsgs = append(errmsgs, msg)
			} else {
				errmsgs = append(errmsgs, fmt.Sprintf("好耶返回了未知的业务错误码: %s", errcode))
			}
		}
		return fmt.Errorf("业务错误: %v", strings.Join(errmsgs, ","))
	}

	return nil
}
