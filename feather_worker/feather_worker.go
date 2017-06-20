package feather_worker

import (
	"github.com/giskook/feather/base"
	"github.com/giskook/feather/conf"
	"github.com/giskook/feather/db_socket"
	"github.com/giskook/feather/http_client"
	"log"
	"time"
)

type FeatherWorker struct {
	ticker     *time.Ticker
	HttpClient *http_client.HttpClient
}

func NewFeatherWorker() *FeatherWorker {
	return &FeatherWorker{
		ticker:     time.NewTicker(time.Duration(conf.GetConf().Http.Interval) * time.Second),
		HttpClient: http_client.NewHttpClient(),
	}
}

func (f *FeatherWorker) Do() {
	for {
		select {
		case <-f.ticker.C:
			f.DoWork()
		}
	}
}

func (f *FeatherWorker) DoWork() {
	start_time := db_socket.GetDBSocket().GetMaxExamTime()
	start_time += 1

	_now := time.Now().Unix()
	log.Printf("start %d , end %d \n", start_time, _now)

	req := http_client.NewHttpRequest(&http_client.HttpRequestParameters{
		Url:           conf.GetConf().Http.Addr,
		ApID:          conf.GetConf().Http.AppID,
		ExamStartTime: start_time,
		ExamEndTime:   _now,
	})

	str_resp := f.HttpClient.Do(req)
	resp, _ := base.ParseResponse(str_resp)
	if resp.Result == true && len(resp.Data) != 0 {
		test_items := db_socket.GetDBSocket().GetAllTestItem()
		db_socket.GetDBSocket().InsertExam(resp, test_items)
	}
}
