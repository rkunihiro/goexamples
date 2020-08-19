package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Todo table row structure
type Todo struct {
	ID     uint64 `json:"id" gorm:"column:id"`
	Title  string `json:"title" gorm:"column:title"`
	Status uint8  `json:"status" gorm:"column:status"`
}

// TableName physical name of Todo table
func (Todo) TableName() string {
	return "Todo"
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
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.Close()
	}()

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	todos := make([]Todo, 0)
	err = db.Order("id ASC").Find(&todos).Error
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(&todos, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
