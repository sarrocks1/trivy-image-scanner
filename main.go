package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

var tpl *template.Template
var outputString string

func main() {
	tpl, _ = tpl.ParseGlob("webpages/*.html")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/imagescan", imageScanHandler)
	http.HandleFunc("/scanresult", scanResultHandler)
	http.ListenAndServe(":8888", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		println(err)
	}
}

func scanResultHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "scanresults.html", outputString)
}

func imageScanHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.FormValue("imageName")
	optionSelected := r.FormValue("outputFormat")
	tpl.ExecuteTemplate(w, "loadpage.html", nil)
	cmd := exec.Command("trivy", "image", "-f", optionSelected, imageName)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		if err.Error() == "exit status 1" {
			outputString = "Please make sure that the image exists"
		} else {
			outputString = err.Error()
		}
	} else {
		outputString = string(output)
	}
}
