package main

import (
	"fmt"
	"log"
	"modulo/infra/mongodb"
	"modulo/router"
	"modulo/utils"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	r := router.Gerar()
	utils.CarregarTemplates() // carregar os templates ParseGlob *.views
	mongodb.InitConnection()  // Inicia o DB
	fmt.Println("Servidor iniciado", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}

}
