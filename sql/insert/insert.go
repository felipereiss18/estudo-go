package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("insert into usuario (nome) value (?)")
	stmt.Exec("Ana")
	stmt.Exec("João")
	res, _ := stmt.Exec("Carlos")

	id, _ := res.LastInsertId()
	fmt.Println("ID =", id)

	linhasAfet, _ := res.RowsAffected()
	fmt.Println("Linhas afetadas =", linhasAfet)
}