package controller

import (
	"html/template"
	"net/http"
	"os"
    "models"
    "time"
)

type HomeController struct {
}

type ParamsHome struct {
    Title string
    ErrMsg string
}
const (
    htmlTitle = "UML META | "
)

func (this *HomeController) HomeAction(w http.ResponseWriter, r *http.Request) {
    models.CheckLogin(w, r)
	tpl_dir := os.Getenv("GOPATH") + "/views/home/"
	t := template.Must(template.ParseFiles(tpl_dir + "home.htm"))
	//准备向页面传参
	p := ParamsHome{Title: htmlTitle + "home"}
	t.Execute(w, p)
}

func (this *HomeController) LoginAction(w http.ResponseWriter, r *http.Request) {
    var errMsg = ""
    if r.Method == "POST" {
        email := r.FormValue("email")
        pass := r.FormValue("pass")
        md5CV := models.Md5Crypt(email, pass, email, "u")
        models.PrintOut(time.Now())
        //已授权用户
        realUsers := models.ParseAuthUser()
        //判断登录用户是否合法
        if models.AuthUser(realUsers, md5CV) == false {
            errMsg = "非授权用户，请确认密码和用户名是否有效"
        } else {
            //验证合法并设置cookie
            cookie := http.Cookie{Name:"_admin_", Value: md5CV, Path: "/", HttpOnly: true, Secure: false, MaxAge: 3600}
            http.SetCookie(w, &cookie)
            http.Redirect(w, r, "/", http.StatusFound)
        }
    } else {

        inter := r.FormValue("do")
        /*
        var doStr string
        switch inter.(type) {
            default:
                doStr = inter.(string)
        }
        */
        if inter == "logout" {
            //models.PrintOut("desk1")
            cookie := http.Cookie{Name:"_admin_", Value: "", Path: "/", HttpOnly: true, Secure: false, MaxAge: -10000}
            http.SetCookie(w, &cookie)
            //http.Redirect(w, r, "/login/", http.StatusFound)
            errMsg = "退出成功"
        }

    }
    tpl_dir := os.Getenv("GOPATH") + "/views/home/"
    t := template.Must(template.ParseFiles(tpl_dir + "login.htm"))
    //准备向页面传参
    p := ParamsHome{Title: htmlTitle + "home/login", ErrMsg: errMsg}
    t.Execute(w, p)
}
