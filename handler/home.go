package handler

import (
	"fmt"
	"net/http"
)

// Home handles homepage requests
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home")
}
