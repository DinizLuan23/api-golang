package main

import (
	"api-go/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o servidor go")
	routes.HandleRequest()
}