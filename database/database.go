package database

import (
	"api/constants"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	conection, err := sql.Open("mysql", constants.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión exitosa.")
	DB = conection
}

func Close() {
	DB.Close()
	fmt.Println("Conexión terminada.")
}

func Ping() {
	if err := DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Sigo conectado")
}

func Exec(query string, args ...any) (sql.Result, error) {
	res, err := DB.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}

	return res, err
}

func Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}
