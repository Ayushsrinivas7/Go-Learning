package main 

import (
"database/sql"
_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB 

func SetUpDB(){
	dsn := "root:@tcp(localhost:3306)/ecommerce"
	var err error 
	db , err = sql.Open("mysql" , dsn)

	if err != nil {
		panic(err)
	}

}

func TeardownDB(){
	if db != nil {
		db.Close()
	}
}
