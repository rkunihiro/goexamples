package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Todo table row structure
type Todo struct {
	ID     uint64 `json:"id"`
	Title  string `json:"title"`
	Status uint8  `json:"status"`
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
	db, err := sql.Open("mysql", dsn)
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
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	todos := make([]Todo, 0)
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}

	jsonData, err := json.MarshalIndent(&todos, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
