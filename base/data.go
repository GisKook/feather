package base

import (
	"encoding/json"
	"log"
)

type HealthData struct {
	ID        string `json:"ID"`
	MachineID string `json:"MachineNo"`
	ExamTime  string `json:"ExamTime"`

	Name           string `json:"PersonName"`
	IDNumber       string `json:"IDNumber"`
	Nation         string `json:"Nation"`
	Sex            string `json:"Sex"`
	Age            int    `json:"Age"`
	Address        string `json:"Address"`
	Telephone      string `json:"Telephone"`
	TestItem       string `json:"JCXM"`
	ChiefComplaint string `json:"ZS"`
	PastHistroy    string `json:"JWS"`
	TestRecord     string `json:"TestData"`
	DiagnoseReport string `json:"ZDBG"`
	Solution       string `json:"ILL"`
}

type PageInfo struct {
	PageIndex int `json:"PageIndex"`
	PageSize  int `json:"PageSize"`
	PageCount int `json:"PageCount"`
}

type Response struct {
	Result   bool          `json:"Success"`
	Message  string        `json:"Message"`
	PageInfo *PageInfo     `json:"PageInfo"`
	Data     []*HealthData `json:"Datas"`
}

func ParseResponse(resp string) (*Response, error) {
	s := new(Response)
	err := json.Unmarshal([]byte(resp), s)
	if err != nil {
		log.Println(err)
	}
	return s, err
}
