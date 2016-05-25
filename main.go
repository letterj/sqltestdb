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
	fmt.Printf("Writing %s records.\n", count)

	db, err := sql.Open("sqlite3", os.Args[1])
	checkErr(err)
	fmt.Printf("Open DB %s.\n", os.Args[1])
	defer db.Close()

	ddl := `
        PRAGMA automatic_index = ON;
        PRAGMA cache_size = 32768;
        PRAGMA cache_spill = OFF;
        PRAGMA foreign_keys = ON;
        PRAGMA journal_size_limit = 67110000;
        PRAGMA locking_mode = NORMAL;
        PRAGMA page_size = 4096;
        PRAGMA recursive_triggers = ON;
        PRAGMA secure_delete = ON;
        PRAGMA synchronous = NORMAL;
        PRAGMA temp_store = MEMORY;
        PRAGMA journal_mode = WAL;
        PRAGMA wal_autocheckpoint = 16384;

				DROP TABLE IF EXISTS test2;

        CREATE TABLE test2 (
            id INT,
          	name TEXT
        );
	`

	_, err = db.Exec(ddl)
	checkErr(err)
	fmt.Println("Process ddl statements")

	stmt, err := db.Prepare("INSERT INTO test2 (id, name) values(?,?)")
	checkErr(err)

	for i := 1; i < count; i++ {
		// insert
		fmt.Println("COUNT", i)
		res, err := stmt.Exec(i, "Testdata")
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
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
