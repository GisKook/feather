package db_socket

import (
	"log"
	"time"
)

const sql_get_max_exam_time string = "SELECT MAX(ExamTime) FROM t_xk_health_machine"

func (db *DbSocket) GetMaxExamTime() int64 {
	stmt, err := db.Db.Prepare(sql_get_max_exam_time)
	defer stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	var max_exam_time string

	err = stmt.QueryRow().Scan(&max_exam_time)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	_time, _ := time.ParseInLocation("2006-01-02 15:04:05", max_exam_time, time.Local)

	return _time.Unix()
}
