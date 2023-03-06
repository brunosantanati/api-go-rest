package main

import (
	"fmt"
	"log"
	"net/http"
)

//Para rodar a aplicação usar: go run main.go
//Acessar http://localhost:8000/

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	HandleRequest()
}
