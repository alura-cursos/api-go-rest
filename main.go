package main

import (
	"fmt"

	"github.com/guilhermeonrails/go-rest-api/models"
	"github.com/guilhermeonrails/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
