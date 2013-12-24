package main

import (
    "controller"
    "net/http"
    "reflect"
    "log"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Path)
    controllerObj := &controller.HomeController{}
    controller := reflect.ValueOf(controllerObj)
    method := controller.MethodByName("HomeAction")
    rv := reflect.ValueOf(r)
    wv := reflect.ValueOf(w)

    method.Call([]reflect.Value{wv, rv})
}

func siteHandler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Path)
    controllerObj := &controller.SiteController{}
    controller := reflect.ValueOf(controllerObj)
    method := controller.MethodByName("SiteAction")
    rv := reflect.ValueOf(r)
    wv := reflect.ValueOf(w)

    method.Call([]reflect.Value{wv, rv})
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Path)
    controllerObj := &controller.AdminController{}
    controller := reflect.ValueOf(controllerObj)
    method := controller.MethodByName("AdminAction")
    rv := reflect.ValueOf(r)
    wv := reflect.ValueOf(w)

    method.Call([]reflect.Value{wv, rv})
}

//登录
func loginHandler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Path)
    controllerObj := &controller.HomeController{}
    controller := reflect.ValueOf(controllerObj)
    method := controller.MethodByName("LoginAction")
    rv := reflect.ValueOf(r)
    wv := reflect.ValueOf(w)

    method.Call([]reflect.Value{wv, rv})
}
