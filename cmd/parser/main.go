package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/parser2.0/internal/db"
	"github.com/parser2.0/internal/parser"
)


func main() {
	database := db.ConnectToDB()
	defer database.Close()
	autos := parser.GetData()
	for _, auto := range autos {
		db.InsertToDB(database, auto)
	}
}
