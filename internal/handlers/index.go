package handlers

import (
	"fmt"
	"net/http"
)

// IndexHandler provides a basic index page with the metrics path in order to
// be found easily. This is a Prometheus good practice.
func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<h1>CSGO Exporter</h1><p><a href='/metrics'>Metrics</a></p>`))
	fmt.Fprint(w)
}
