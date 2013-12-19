package controller

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type HomeController struct {
}

type ParamsHome struct {
    Title string
}

func (this *HomeController) HomeAction(w http.ResponseWriter, r *http.Request) {
    log.Println("home")
	tpl_dir := os.Getenv("GOPATH") + "/views/home/"
	t := template.Must(template.ParseFiles(tpl_dir + "home.htm"))
	//准备向页面传参
	p := ParamsHome{Title: "home"}
	t.Execute(w, p)
}
