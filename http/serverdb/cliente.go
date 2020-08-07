package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Usuario struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
}

func UsuarioHandler(resp http.ResponseWriter, req *http.Request)  {
	sid := strings.TrimPrefix(req.URL.Path, "/usuarios/")

	id, _ := strconv.Atoi(sid)

	switch {
	case req.Method == "GET" && id > 0:
		usuarioPorID(resp, req, id)
	case req.Method == "GET":
		usuarioTodos(resp, req)
	default:
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "Desculpa...")
	}
}

func usuarioTodos(resp http.ResponseWriter, req *http.Request) {
	db := conexaoDB()
	defer db.Close()

	var lista []Usuario

	rows, _ := db.Query("select id, nome from usuario")

	defer rows.Close()

	for rows.Next(){
		var usu Usuario
		rows.Scan(&usu.ID, &usu.Nome)

		lista = append(lista, usu)
	}

	json, _ := json2.Marshal(lista)

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprint(resp, string(json))
}

func usuarioPorID(resp http.ResponseWriter, req *http.Request, id int) {
	db := conexaoDB()
	defer db.Close()

	var usu Usuario
	db.QueryRow("select id, nome from usuario where id = ?", id).Scan(&usu.ID, &usu.Nome)

	json, _ := json2.Marshal(usu)

	resp.Header().Set("Content-Type", "application/json")
	fmt.Fprint(resp, string(json))
}

func conexaoDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil{
		log.Fatal(err)
	}

	return db
}