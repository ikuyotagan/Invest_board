package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "jwt"
)

var DB *sql.DB

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	DB = db
	fmt.Println("Connection successfull")
}
