package main

import (
	"fmt"

	"github.com/brunosantanati/api-go-rest/database"
	"github.com/brunosantanati/api-go-rest/routes"
)

//Para rodar a aplicação usar: go run main.go
//Acessar http://localhost:8000/

func main() {
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
