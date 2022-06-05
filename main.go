package main

import (
	"fmt"

	"github.com/igorcavalcantedb/golang-rest-crud/database"
	"github.com/igorcavalcantedb/golang-rest-crud/models"
	"github.com/igorcavalcantedb/golang-rest-crud/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com GO")

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
