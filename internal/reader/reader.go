package internal

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://dou.ua/lenta/tags/RSS/feed/"

func ReadRss() string {

	return string(ReadRssBytes())

}

func ReadRssBytes() []byte {

	resp, err := getResponse()
	if err != nil {
		panic(err)
	}

	body, err := parseResponse(resp)
	if err != nil {
		panic(err)
	}

	return body
}

func getResponse() (*http.Response, error) {

	resp, err := http.Get(Url)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Errorf("Status Code is %d, but should be 200", resp.StatusCode)
	}

	return resp, nil
}

func parseResponse(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp.Body.Close()

	return body, nil
}
