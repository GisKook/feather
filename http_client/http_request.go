package http_client

import (
	"github.com/giskook/feather/base"
	"github.com/giskook/feather/conf"
	"net/http"
	"net/url"
)

const (
	APID            string = "apid"
	EXAM_START_TIME string = "ExamStartTime"
	EXAM_END_TIME   string = "ExamEndTime"
	PAGE_INDEX      string = "PageIndex"
	PAGE_SIZE       string = "PageSize"
)

type HttpRequestParameters struct {
	Url           string
	ApID          string
	ExamStartTime int64
	ExamEndTime   int64
	PageIndex     int
	PageSize      int
}

type HttpRequest struct {
	Request *http.Request
}

func NewHttpRequest(parameters *HttpRequestParameters) *HttpRequestParameters {
	req, err := http.NewRequest("POST", parameters.Url, nil)
	if err != nil {
		base.PrintError(err)

		return nil
	}
	form := url.Values{}
	form.Add(APID, parameters.ApID)
	form.Add(EXAM_START_TIME, parameters.ExamStartTime)
	form.Add(EXAM_END_TIME, parameters.ExamEndTime)
	form.Add(PAGE_INDEX, parameters.PageIndex)
	form.Add(PAGE_SIZE, parameters.PageSize)

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return &HttpRequest{
		Request: req,
	}
}
