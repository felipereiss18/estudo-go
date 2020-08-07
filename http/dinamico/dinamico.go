package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(resp http.ResponseWriter, req *http.Request)  {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(resp, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
