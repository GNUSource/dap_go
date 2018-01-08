package service

import (
	"conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"module/function"
	"module/logger"
	"strings"
)

func HandleRawData() {
	db, err := sql.Open("mysql", conf.GetDsn())
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	sql := `SELECT cd, wl_name
			FROM raw_sap_yy
			where wl_spec is null
				and wl_name is not null
			order by cd asc;`
	rows, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	var (
		cd     string
		o_name string
	)

	for rows.Next() {
		err = rows.Scan(&cd, &o_name)
		if err != nil {
			log.Fatal(err)
		}
		logger.Trace.Printf("origin str is %v\n", o_name)

		// 先将 name 中需要剔除的字符串剔除
		deleteStr := "(通说明书)|说明书|(政府)|政府|标盒|中盒|标签|铝箔|大箱|(通用)|(西班牙文)|(不干胶)|(质量和疗效一致性评价)|(自销)|(膜厚)|(通用)"
		p1_name := function.DeleteSubStr(o_name, deleteStr)
		logger.Trace.Printf("p1_name is %v\n", p1_name)
		// 将 p1_name 中的规格信息提取出来
		regexpRule := `\(\S+\)|\d+(.\d+)?(/)?(mg|g|ml|mm|粒|盒|袋|片|板)(/|\*)?\d+(.\d+)?(mg|g|ml|mm|粒|盒|袋|片|板)|\d+(.\d+)?(mg|g|ml|mm|粒|盒|袋|片|板)`
		p2_name, matched_str := function.DrawStr(p1_name, regexpRule)
		logger.Trace.Printf("p2_name is %v\n", p2_name)

		var spec string
		spec = strings.Join(matched_str, "||")
		spec = function.DeleteSubStr(spec, "(|)")
		logger.Trace.Printf("spec is %v\n", spec)

		sql := "update raw_sap_yy set wl_name = ? , wl_spec = ? where cd = ?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(p2_name, spec, cd)
		if err != nil {
			log.Fatal(err)
		}
		stmt.Close()
	}

}
