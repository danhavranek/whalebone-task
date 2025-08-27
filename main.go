package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danhavranek/whalebone-task/database"
	"github.com/danhavranek/whalebone-task/routes"
)

const port uint16 = 8090

func main() {
	routes.Init()

	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	portString := fmt.Sprintf(":%d", port)
	log.Printf("Server running on port [%s]", portString)
	http.ListenAndServe(portString, nil)
}
