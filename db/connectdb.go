package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connectdb() *sql.DB {
	fmt.Println("Go MySql Connecting...")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/school")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
