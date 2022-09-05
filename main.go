package main

import (
	"fmt"

	"github.com/hallex-abreu/api-go-rest/database"
	"github.com/hallex-abreu/api-go-rest/routes"
)

func main() {
	database.ConnectaComBancoDeDados()
	fmt.Println("Iniciando o servidor rest com go")
	routes.HandleRequest()
}
