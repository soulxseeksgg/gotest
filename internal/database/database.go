package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConfig() {

	db, err := sql.Open("mysql", "root:rootpassword@tcp(127.0.0.1:3306)/backend")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("error database connection :", err)
	} else {
		fmt.Println("database connection successfully")
	}

	rows, err := db.Query("SELECT name, email FROM tb_user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, email string
		if err := rows.Scan(&name, &email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name: %s, email: %s\n", name, email)
	}
}
