package internal

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://dou.ua/lenta/tags/RSS/feed/"

// ReadRss fetches the RSS feed from the configured URL and returns it as a string.
func ReadRss() string {

	return string(ReadRssBytes())

}

// ReadRssBytes fetches the RSS feed from the configured URL and returns its raw content as a byte slice.
func ReadRssBytes() []byte {

	resp := getResponse()

	body := parseResponse(resp)

	return body
}

// getResponse performs an HTTP GET request to the RSS feed URL and returns the HTTP response.
// It prints an error message if the request fails or the status code is not 200.
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

// parseResponse reads the body of the given HTTP response and returns it as a byte slice.
// It closes the response body after reading and prints an error message if reading fails.
func parseResponse(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()

	return body
}
