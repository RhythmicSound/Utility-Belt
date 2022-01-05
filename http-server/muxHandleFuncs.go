package httpserver

import (
	"fmt"
	"net/http"
)

//HealthcheckHandler is a quick mux route useful only for responding to checks of
// the server being alive
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "200 OK")
}
