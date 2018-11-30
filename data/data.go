package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//const connstr = "user=postgres dbname=admin sslmode=disable"
const connstr = "postgres://postgres:admin@localhost/admin?sslmode=disable"

func Insert(person []byte) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
	var cols string
	output := db.QueryRow(`INSERT INTO person(ID,Firstname,Lastname,Address)

	VALUES(person.ID,person.Firstname,person.Lastname,person.Address) RETURNING ID,Firstname,Lastname,Address`).Scan(&cols)
        fmt.Println("this is test")
	fmt.Println(output)
}
func DisplayAll() {
	db, err := sql.Open("postgres", connstr)

	if err != nil {
		log.Fatal(err)
	}
//	_, err = db.Query("select * from person")
	rows, err := db.Query("SELECT * FROM admin")
	fmt.Println(rows)
}
