package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	trans, _ := db.Begin()

	stmt, _ := trans.Prepare("insert into usuario (id, nome) values (?, ?)")

	stmt.Exec(2000, "Paulo")
	stmt.Exec(2001, "Clara")
	_, err = stmt.Exec(1, "Ítalo")

	if err != nil {
		trans.Rollback()
		log.Fatal(err)
	}

	trans.Commit()
}
