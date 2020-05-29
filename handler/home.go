package handler

import (
	"html/template"
	"net/http"
)

type pageData struct {
	Result string
}

// Home handles homepage requests
func Home(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("./template/home.gohtml"))

	data := struct {
		Result string
	}{
		Result: "",
	}

	err := html.Execute(w, data)
	if err != nil {
		http.Error(w, "there was a problem serving the template", http.StatusInternalServerError)
	}
}
