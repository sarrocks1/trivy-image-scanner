package main

import (
	"html/template"
	"net/http"
	"os/exec"
)

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("webpages/*.html")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/imagescan", imageScanHandler)
	http.ListenAndServe(":8888", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		println(err)
	}
}

func imageScanHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.FormValue("imageName")
	println(imageName)
	tpl.ExecuteTemplate(w, "loadpage.html", nil)
	cmd := exec.Command("trivy image -f json " + imageName)
	output, err := cmd.Output()
	if err != nil {
		println("Error:", err)
	}
	println(output)
}
