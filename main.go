package main

import (
	"fmt"

	"github.com/brunosantanati/api-go-rest/models"
	"github.com/brunosantanati/api-go-rest/routes"
)

//Para rodar a aplicação usar: go run main.go
//Acessar http://localhost:8000/

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
