package main

import (
	"net/http"

	"github.com/danhavranek/whalebone-task/database"
	"github.com/danhavranek/whalebone-task/routes"
)

func main() {
	routes.Init()

	if err := database.Init(); err != nil {
		panic(err)
	}

	http.ListenAndServe(":8090", nil)
}
