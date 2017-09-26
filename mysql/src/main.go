package main

import (
	"database/sql"
	"fmt"
	"time"

	"net"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:3306"
	dbusername = "cza"
	dbpassword = "123"
	dbname     = "czatest1"
)

func main() {
	connUri := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbusername, dbpassword, dbhostsip, dbname)

	fmt.Println(connUri)
	db, err := sql.Open("mysql", connUri)
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt, err := db.Prepare("INSERT testtable SET username=?,departname=?,created=?")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := stmt.Exec("cza", "ccdd", time.Now().String())
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("id =", id)
}
