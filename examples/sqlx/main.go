package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Todo table row structure
type Todo struct {
	ID     uint64 `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Status uint8  `json:"status" db:"status"`
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	dsn := "user:pass@tcp(0.0.0.0:3306)/dbname?parseTime=true&loc=Asia%2FTokyo&rejectReadOnly=true"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sql := "SELECT id,title,status FROM Todo ORDER BY id"
	stmt, err := db.PrepareNamed(sql)
	if err != nil {
		panic(err)
	}

	todos := make([]Todo, 0)
	err = stmt.Select(&todos, map[string]interface{}{})
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(&todos, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
