package oumlagosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	session http.Client
	clientTransport transport
}

// type Config struct {
// 	BaseURL string
// 	ApiKey string
// 	Env	string
// }

func NewClient(baseURL string, apiKey string, headers ...map[string]string) (*Client, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	t := transport{
		header: http.Header{},
		baseURL: *url,
	}

	t.header.Add("Content-Type", "application/json")
	t.header.Add("x-api-key", fmt.Sprintf("Bearer %s", apiKey))

	c := Client{
		session: http.Client{Transport: t},
		clientTransport: t,
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			c.clientTransport.header.Set(key, value)
		}
	}

	return &c, nil
}

// NewRequest will create new request with method, path and body
// If body is not nil, it will be marshalled into json
func (c *Client) NewRequest(method, path string, body ...interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if len(body) > 0 && body[0] != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body[0])
		if err != nil {
			return nil, err
		}
	}
	url := c.clientTransport.baseURL.String() + path
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Do will send request using the c.sessionon which it is called
// If response contains body, it will be unmarshalled into v
// If response has err, it will be returned
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.session.Do(req)
	if err != nil {
		return nil, err
	}

	err = checkForError(resp)
	if err != nil {
		return resp, err
	}

	if resp.Body != nil && v != nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp, err
		}
		err = json.Unmarshal(body, &v)
		fmt.Println(string(body))
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func checkForError(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c < 400 {
		return nil
	}

	errorResponse := &OumlaError{
		Status: resp.StatusCode,
		Path: resp.Request.URL.Path,
	}

	data, err := io.ReadAll(resp.Body)
	fmt.Println(string(data))
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}
	errorResponse.SetErrorType(resp.StatusCode)

	return errorResponse
}

type transport struct {
	header http.Header
	baseURL url.URL
}

func (t transport) RoundTrip(request *http.Request) (*http.Response, error) {
	for headerName, values := range t.header {
		for _, val := range values {
			request.Header.Add(headerName, val)
		}
	}
	request.URL = t.baseURL.ResolveReference(request.URL)
	return http.DefaultTransport.RoundTrip(request)
}