package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		dsn = "root:root@tcp(127.0.0.1:3306)/crud_db"
	}

	var db *sql.DB
	var err error

	// Tenta conectar 5 vezes antes de desistir
	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Conectado ao MySQL com sucesso!")
				return db, nil
			}
		}

		fmt.Printf("Aguardando banco de dados... tentativa %d\n", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, err
}
