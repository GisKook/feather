package db_socket

import (
	"fmt"
	"github.com/giskook/feather/base"
	"log"
)

const sql_get_test_item string = "SELECT id  FROM t_xk_jcxm where JCXM = ?"

func (db *DbSocket) GetTestItem(test_item string) int {
	log.Println(sql_get_test_item)
	stmt, err := db.Db.Prepare(sql_get_test_item)
	defer stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	var id int

	err = stmt.QueryRow(test_item).Scan(&id)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return id
}

const sql_set_test_item string = "INSERT INTO t_xk_jcxm (JCXM) values(%s)"

func (db *DbSocket) SetTestItem(test_item string) {
	_sql := fmt.Sprintf(sql_set_test_item, test_item)

	log.Println(_sql)
	_, err := db.Db.Exec(_sql)
	if err != nil {
		log.Println(err.Error())
	}
}

const sql_get_all_test_item string = "select JCXM, id  from t_xk_jcxm"

func (db *DbSocket) GetAllTestItem() map[string]int {
	stmt, err := db.Db.Prepare(sql_get_all_test_item)
	defer stmt.Close()
	log.Println(sql_get_all_test_item)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	var test_item string
	var id int

	rows, er := stmt.Query()
	base.PrintError(er)
	defer rows.Close()

	test_items := make(map[string]int)
	for rows.Next() {
		if e := rows.Scan(&test_item, &id); e != nil {
			base.PrintError(e)
		}
		test_items[test_item] = id
	}

	return test_items
}
