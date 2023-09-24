package v1

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB
var err error

func init()  {
	DB, err =sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err !=nil{
		log.Fatal(err)
	}
	DB.SetConnMaxLifetime(50*time.Second)
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(50)
}



