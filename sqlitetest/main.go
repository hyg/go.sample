// sqlitetest project main.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.s3db")
	checkErr(err)

	stmt, err := db.Prepare("insert into Carrier values (?,?,?,?,?,?,?,?,?,?)")
	checkErr(err)

	//insert into carrier values (1,2,3,'万家康',1,'一家专业冷链物流公司',now(),1,null,null),(2,2,3,'康奇朋友',1,'来自CSA农场的物流合作社',now(),1,null,null);
	/*
		res, err := stmt.Exec(1, 2, 3, "万家康", 1, "一家专业冷链物流公司", time.Now(), 1, nil, nil)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Printf("CarrierID:%d\n", id)
	*/
	res, err := stmt.Exec(2, 2, 3, "康奇朋友", 1, "来自CSA农场的物流合作社", time.Now(), 1, nil, nil)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Printf("CarrierID:%d\n", id)

	rows, err := db.Query("SELECT CarrierID ,Name  from Carrier")
	checkErr(err)

	for rows.Next() {
		var ID int
		var Name string

		err = rows.Scan(&ID, &Name)
		checkErr(err)

		fmt.Printf("ID=%d  Name=%s\n", ID, Name)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
