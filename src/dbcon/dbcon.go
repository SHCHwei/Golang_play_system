package dbcon

import (

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func DBT(){
	
	dbSettings := fmt.Sprintf("%s:%s@(%s)/%s","root","","192.168.99.137","gp")
	DB, err = sql.Open("mysql",dbSettings)
		
	if err != nil{
		panic(err)
	} else {
		rows, err := DB.Query("Select * from user where id = 0")

		if err != nil{
			panic(err)
		}
	
		for rows.Next() {
			var user string
			err = rows.Scan(&user)
			if err != nil{
				panic(err)
			}
	
			return 
		}
	}
}


func Select(sqlString string){

	rows, err := DB.Query(sqlString)

	if err != nil{
		panic(err)
	}

	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil{
			panic(err)
		}

		return
	}
}
