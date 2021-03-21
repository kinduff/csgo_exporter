package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	response := `<h1>CSGO Exporter</h1><p><a href='/metrics'>Metrics</a></p>`
	fmt.Fprintf(w, response)
}
