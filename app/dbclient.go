package app

import (
	"fmt"
	"os"
	"payments/logs"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func dbclient() *sqlx.DB {
	usr := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWD")
	addr := os.Getenv("DB_ADDR")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", usr, pass, addr, port, name)
	client, err := sqlx.Open("mysql", dsn)
	if err != nil {
		logs.Fatal(err.Error())
	}

	//pool settings:
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
