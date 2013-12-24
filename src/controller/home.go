package controller

import (
	"html/template"
	"log"
	"net/http"
	"os"
    "models"
)

type HomeController struct {
}

type ParamsHome struct {
    Title string
}
const (
    htmlTitle = "UML META | "
)

func (this *HomeController) HomeAction(w http.ResponseWriter, r *http.Request) {
    log.Println("home")
    models.CheckLogin(w, r)
	tpl_dir := os.Getenv("GOPATH") + "/views/home/"
	t := template.Must(template.ParseFiles(tpl_dir + "home.htm"))
	//准备向页面传参
	p := ParamsHome{Title: htmlTitle + "home"}
	t.Execute(w, p)
}

func (this *HomeController) LoginAction(w http.ResponseWriter, r *http.Request) {
    log.Println("home/login")
    if r.Method == "POST" {
        email := r.FormValue("email")
        pass := r.FormValue("pass")
        log.Println("pass:", pass)
        cookie := http.Cookie{Name:"_admin_", Value: email, Path: "/", HttpOnly: true, Secure: true, MaxAge: 1000}
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/", http.StatusFound)
    }
	tpl_dir := os.Getenv("GOPATH") + "/views/home/"
	t := template.Must(template.ParseFiles(tpl_dir + "login.htm"))
	//准备向页面传参
	p := ParamsHome{Title: htmlTitle + "home/login"}
	t.Execute(w, p)
}
