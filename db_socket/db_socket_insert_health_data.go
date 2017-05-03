package db_socket

import (
	"fmt"
	"github.com/giskook/feather/base"
	//"log"
)

const (
	TRANS_TABLE_NAME_FMT    string = "t_xk_health_machine"
	SQL_INSERT_HEALTH_TABLE string = "INSERT %s (MachineID, ExamTime, JCXMID, PersonName, IDNumber, Nation, Sex, Age, Address, Telephone, ZS, JWS, TestData, ZDBG, JJFA, ExamUuid) VALUES('%s', '%s', %d, '%s', '%s', '%s','%s', %d, '%s', '%s','%s', '%s','%s','%s','%s','%s')"
)

func (socket *DbSocket) InsertExam(resp *base.Response, test_items map[string]int) {
	tx, err := socket.Db.Begin()
	base.PrintError(err)
	for _, v := range resp.Data {
		if v.IDNumber != "" {
			_sql := FmtSQL(v, test_items)
			//log.Println(_sql)
			_, err = tx.Exec(_sql)
			base.PrintError(err)
		}
	}
	err = tx.Commit()
	base.PrintError(err)
}

func FmtSQL(data *base.HealthData, test_items map[string]int) string {

	insert_sql := fmt.Sprintf(SQL_INSERT_HEALTH_TABLE, TRANS_TABLE_NAME_FMT, data.MachineID, data.ExamTime, test_items[data.TestItem], data.Name, data.IDNumber, data.Nation, data.Sex, data.Age, data.Address, data.Telephone, data.ChiefComplaint, data.PastHistroy, data.TestRecord, data.DiagnoseReport, data.Solution, data.ID)

	return insert_sql
}
