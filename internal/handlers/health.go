package handlers

import (
	"fmt"
	"net/http"
)

// HealthHandler provides a plain OK with an HTTP 200 code to
// prove that the server is running and responding.
func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	fmt.Fprint(w)
}
