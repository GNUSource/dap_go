package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"conf"
	"log"
)

func main() {
	db, err := sql.Open("mysql", conf.GetDsn())
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	rows, err := db.Query("SELECT wl_des, wl_name, wl_spec FROM raw_sap_yy limit 5")
	if err != nil {
		log.Fatal(err)
	}

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		//record := make(map[string]string)
		for i, col := range values {
			fmt.Printf("i is %s, col is %s\n", columns[i], string(col.([]byte)))
		}
		//fmt.Println(record)
	}


}
