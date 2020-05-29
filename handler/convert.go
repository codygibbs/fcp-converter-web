package handler

import (
	"encoding/xml"
	"html/template"
	"net/http"

	converter "github.com/codygibbs/fcp-converter"
)

// Convert handles an XML file conversion request
func Convert(w http.ResponseWriter, r *http.Request) {
	html := template.Must(template.ParseFiles("./template/home.gohtml"))

	source := r.FormValue("import")

	xeml := converter.ImportRawXEML([]byte(source))

	cb, err := xml.MarshalIndent(xeml, "", "  ")
	if err != nil {
		http.Error(w, "there was a problem serving the template", http.StatusInternalServerError)
	}

	data := pageData{string(cb)}

	err = html.Execute(w, data)
	if err != nil {
		http.Error(w, "there was a problem serving the template", http.StatusInternalServerError)
	}
}
