package http_client

import (
	"github.com/giskook/feather/base"
	"net/http"
)

type HttpClient struct {
	Instance *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Instance: &http.Client{},
	}
}

func (c *HttpClient) Do(req *HttpRequest) {
	resp, err := c.Do(req.Request)
	if err != nil {
		base.PrintError(err)
		return ""
	}
	defer resp.Body.Close()
}
