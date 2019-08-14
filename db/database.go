package db

import (
	"database/sql"
	"fmt"
	//"log"
	//"os"

	_ "go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	serverName := "localhost:3306"
	user := "root"
	password := ""
	dbName := "go"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
