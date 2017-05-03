package http_client

import (
	"github.com/giskook/feather/base"
	"io/ioutil"
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

func (c *HttpClient) Do(req *HttpRequest) string {
	resp, err := c.Instance.Do(req.Request)
	if err != nil {
		base.PrintError(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
