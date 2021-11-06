package database

import (
	"database/sql"
	"fmt"
	"user-api/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver
)

func Connect() (*sql.DB, error) {
	fmt.Print(config.StringConnectionDatabase)
	db, err := sql.Open("mysql", config.StringConnectionDatabase)
	if err != nil {
		fmt.Print("Erro open")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		fmt.Print("Erro ping")

		db.Close()
		return nil, err
	}

	return db, nil
}
