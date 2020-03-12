package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	dname := "user:1111@tcp(127.0.0.1:3306)/po?charset=utf8&loc=local"
	db, err := sql.Open("mysql", dname)
	if err != nil {
		fmt.Println(err)
	}
	//设置连接池
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(7 * time.Hour)
	fmt.Println(db.Ping())
	defer db.Close()
}
