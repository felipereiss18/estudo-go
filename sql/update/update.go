package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("update usuario set nome = ? where id = ?")
	stmt.Exec("Karol", 1)
	stmt.Exec("Alexandre", 2)

	stmt2, _ := db.Prepare("delete from usuario where id = ?")
	stmt2.Exec(3)
}
