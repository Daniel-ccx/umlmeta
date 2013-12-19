package controller

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type AdminController struct {
}

type ParamsAdmin struct {
	Title string
}

func (this *AdminController) AdminAction(w http.ResponseWriter, r *http.Request) {
	log.Println("admin")
	tpl_dir := os.Getenv("GOPATH") + "/views/admin/"
	t := template.Must(template.ParseFiles(tpl_dir + "home.htm"))
	//准备向页面传参
	p := ParamsAdmin{Title: "admin"}
	t.Execute(w, p)
}
