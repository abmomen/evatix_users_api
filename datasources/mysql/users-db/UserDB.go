package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlUsersUserName = "mysqlUsersUserName"
	mysqlUsersPassword = "mysqlUsersPassword"
	mysqlUsersHost = "mysqlUsersHost"
	mysqlUsersSchema = "mysqlUsersSchema"

)

var (
	Client *sql.DB

	userName = os.Getenv(mysqlUsersUserName)
	password = os.Getenv(mysqlUsersPassword)
	host 	 = os.Getenv(mysqlUsersHost)
	schema 	 = os.Getenv(mysqlUsersSchema)
)


func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", userName, password, host, schema)
	var err error

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil{
		panic(err)
	}
	log.Println("database successfully configured")
}