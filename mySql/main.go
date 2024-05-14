package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB()(err error) {
	dsn := "root:adminadmin@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed, err:%v\n", err)
		return
	}
	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?, department=?, created=?")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	id, err:= stmt.Exec("wuda", "啥也不会部", "2045-12-12")
	if err != nil {
		return
	}
	fmt.Print(id)
	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	if err != nil {
		return
	}
	res, err := stmt.Exec("astaxieupdate", 5)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	affect, err := res.RowsAffected()
	fmt.Print(affect)
	// 查询数据
	rows, err := db.Query("select * from userinfo")
	if err != nil {
		return
	}

	for rows.Next() {
		var uid int
		var username string
		var department string 
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		if err != nil {
			return
		}
		fmt.Println(uid, username, department, created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	res, err = stmt.Exec(5)

	defer db.Close()
}	