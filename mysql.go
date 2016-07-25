package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "sa:abcd-1234@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	//获取时间戳
	timestamp := time.Now().Unix()
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")

	res, err := stmt.Exec("astaxie", "研发部门", tm)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	fmt.Println(scanArgs)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	/*	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}*/
	/*
		//删除数据
		stmt, err = db.Prepare("delete from userinfo where uid=?")
		checkErr(err)
		res, err = stmt.Exec(id)
		checkErr(err)
		affect, err = res.RowsAffected()
		checkErr(err)
		fmt.Println(affect)*/
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
