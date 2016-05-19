package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	count, err := strconv.Atoi(os.Args[2])
	checkErr(err)

	db, err := sql.Open("sqlite3", os.Args[1])
	checkErr(err)
	fmt.Println("Open DB")

	stmt, err := db.Prepare("DROP TABLE IF EXISTS test2;")
	checkErr(err)
	fmt.Println("Drop table if exists")

	res, err := stmt.Exec()
	checkErr(err)

	stmt, err = db.Prepare("CREATE TABLE test2 (id int, name text);")
	checkErr(err)
	fmt.Println("Create table")

	res, err = stmt.Exec()
	checkErr(err)

	stmt, err = db.Prepare("INSERT INTO test2 (id, name) values(?,?)")
	checkErr(err)

	for i := 1; i < count; i++ {
		// insert
		fmt.Println("COUNT", i)
		res, err = stmt.Exec(i, "Testdata")
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println("LAST_INSERTED_ID", id)

		affect, err := res.RowsAffected()
		checkErr(err)
		fmt.Println("ROWS_AFFECTED", affect)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("QUERYING THE DB")
	fmt.Println("===============")
	// query
	rows, err := db.Query("SELECT * FROM test2")
	checkErr(err)

	var test2_ID int
	var test2_Name string

	for rows.Next() {
		err = rows.Scan(&test2_ID, &test2_Name)
		checkErr(err)
		fmt.Printf("ID: %d, NAME: %s\n", test2_ID, test2_Name)
	}

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
