package main

import (
	"api-go/database"
	"api-go/models"
	"api-go/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	database.ConectaBd()
	
	fmt.Println("Iniciando o servidor go")
	routes.HandleRequest()
}
