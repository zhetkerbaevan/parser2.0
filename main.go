package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/parser2.0/act"
	"github.com/parser2.0/entity"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/parser_go")
	fmt.Println("connected to DB")
	checkerr(err)
	return db
}

func main() {
	db := connectToDB()
	defer db.Close()
	autos := act.GetData()
	for _, auto := range autos {
		insertToDB(db, auto)
	}
}

func insertToDB(db *sql.DB, auto entity.Automobile) {
	db.Exec("INSERT INTO `automobiles` (`id`, `model`, `year`, `price`) VALUES (?, ?, ?, ?)", auto.ID, auto.Model, auto.Year, auto.Price)
}
