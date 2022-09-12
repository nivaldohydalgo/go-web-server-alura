package main

import (
	"net/http"
	"nivaldo/web-server/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
