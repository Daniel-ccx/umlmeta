package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//向页面输出的参数
type ParamsSite struct {
    Title string
}

type SiteController struct {
}

func (this *SiteController) SiteAction(w http.ResponseWriter, r *http.Request) {
	tpl_dir := os.Getenv("GOPATH") + "/views/site/"
	fmt.Println("abdfasf")
	t := template.Must(template.ParseFiles(tpl_dir + "home.htm"))
	//解析参数
	r.ParseForm()
    p := ParamsSite{Title: "site"}
	//准备向页面传参
	t.Execute(w, p)
}
