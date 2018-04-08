package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type row struct {
	name string
	text string
}

type user struct {
	dbuser, dbpass string
}

func setDbUser() user {
	dbuser := os.Getenv("DBUser")
	dbpass := os.Getenv("DBPass")
	return user{dbuser, dbpass}
}
func toDb(name, text string) {
	user := setDbUser()
	db, err := sql.Open("mysql", user.dbuser+":"+user.dbpass+"@/simplechatlog")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	sqlStr := "insert into `log` values('" + name + "','" + text + "');"
	if _, err := db.Exec(sqlStr); err != nil {
		log.Fatal(err)
	}

	time := time.Now()
	fmt.Println(time, name, text)
}

func fromDb() []string {
	user := setDbUser()
	db, err := sql.Open("mysql", user.dbuser+":"+user.dbpass+"@/simplechatlog")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from log")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	rowsSl := []string{}
	for rows.Next() {
		var name, text string
		if err := rows.Scan(&name, &text); err != nil {
			log.Fatalln(err)
		}
		rowsSl = append(rowsSl, name+":「"+text+"」")
	}
	return rowsSl
}
