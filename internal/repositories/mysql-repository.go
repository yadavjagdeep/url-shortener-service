package repositories

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlClient *sql.DB
	dbName      string
)

func InitMysqlClient(uri string, database string) {
	url := uri + "/" + database + "?parseTime=true"
	log.Printf("Url: %v", url)

	client, err := sql.Open("mysql", url)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	mysqlClient = client
	dbName = database

}

func GetMySqlClient() *sql.DB {
	return mysqlClient
}

func GetDBName() string {
	return dbName
}
