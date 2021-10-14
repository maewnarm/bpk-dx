package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// create file dblogin.go and write below code

// package models

// const (
// 	DbDatabase = "db name"
// 	DbUser     = "db user"
// 	DbPassword = "db password"
// )

func connect() (*sqlx.DB, error) {
	DBurl := fmt.Sprintf("postgres://postgres:postgres@bpk-pg:5432/%s?sslmode=disable", DbDatabase)
	//DBurl := fmt.Sprintf("postgres://0.0.0.0:5432/%s?sslmode=disable&user=%s&password=%s", DbDatabase, DbUser, DbPassword)
	return sqlx.Connect("postgres", DBurl)
}
