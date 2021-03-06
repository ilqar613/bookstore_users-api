package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username" //=root
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host" //127.0.0.1:3306
	mysql_users_schema   = "mysql_users_schema" //users_db
)

var (
	Client   *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	//log.Println(fmt.Sprintf("connection params %s",dataSourceName))

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
