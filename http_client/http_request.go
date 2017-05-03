package http_client

import (
	"bytes"
	"fmt"
	"github.com/giskook/feather/base"
	"net/http"
	"strconv"
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

func NewHttpRequest(parameters *HttpRequestParameters) *HttpRequest {
	start_time := strconv.FormatInt(parameters.ExamStartTime, 10)
	end_time := strconv.FormatInt(parameters.ExamEndTime, 10)
	params := fmt.Sprintf("%s=%s&%s=%s&%s=%s", APID, parameters.ApID, EXAM_START_TIME, start_time, EXAM_END_TIME, end_time)
	req, err := http.NewRequest("POST", parameters.Url, bytes.NewBuffer([]byte(params)))
	if err != nil {
		base.PrintError(err)

		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return &HttpRequest{
		Request: req,
	}
}
