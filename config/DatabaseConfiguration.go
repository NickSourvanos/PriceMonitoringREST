package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "ProductsDb"
)

func GetDatabaseConnection() (*sql.DB, error) {

	desc := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := createConection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConection(desc string)(*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil{
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}

//var DB *sql.DB
//
//func InitiateDatabase() {
//	config := databaseConfiguration()
//	var err error
//
//	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config[DBHOST],
//		config[DBPORT], config[DBUSER], config[DBPASS], config[DBNAME])
//
//	DB, err = sql.Open("postgres", psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//
//	err = DB.Ping()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Successfully connected!")
//}
//
//func CloseDatabaseConnection() {
//	DB.Close()
//}
//
//func databaseConfiguration() map[string]string{
//	conf := make(map[string]string)
//	conf[DBHOST] = "localhost"
//	conf[DBPORT] = "5432"
//	conf[DBUSER] = "postgres"
//	conf[DBPASS] = "root"
//	conf[DBNAME] = "usersdb"
//	return conf
//}