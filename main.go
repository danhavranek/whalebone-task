package main

import (
	"net/http"

	"github.com/danhavranek/whalebone-task/database"
	"github.com/danhavranek/whalebone-task/routes"
)

func main() {
	routes.Init()

	err := database.Init()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8090", nil)
}
