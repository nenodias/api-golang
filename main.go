package main

import (
	"github.com/nenodias/api-golang/models"
	"github.com/nenodias/api-golang/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	routes.HandleRequest()
}
