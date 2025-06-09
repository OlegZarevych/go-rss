package internal

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://dou.ua/lenta/tags/RSS/feed/"

func ReadRss() string {

	resp := getResponse()

	body := parseResponse(resp)

	return string(body)

}

func ReadRssBytes() []byte {

	resp := getResponse()

	body := parseResponse(resp)

	return body
}

func getResponse() *http.Response {

	resp, err := http.Get(Url)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 200 {
		fmt.Errorf("Status Code is %d, but should be 200", resp.StatusCode)
	}

	return resp
}

func parseResponse(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()

	return body
}
