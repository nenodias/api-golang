package main

import (
	"github.com/nenodias/api-golang/database"
	"github.com/nenodias/api-golang/models"
	"github.com/nenodias/api-golang/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
