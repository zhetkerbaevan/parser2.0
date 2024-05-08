package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/parser2.0/internal/model"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/parser_go")
	fmt.Println("connected to DB")
	checkerr(err)
	return db
}

func InsertToDB(db *sql.DB, auto model.Automobile) {
	db.Exec("INSERT INTO `automobiles` (`id`, `model`, `year`, `price`) VALUES (?, ?, ?, ?)", auto.ID, auto.Model, auto.Year, auto.Price)
}
