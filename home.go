package main

import (
	"log"
	"net/http"
)

func main() {
	//var cfg = models.ParseCfg("mysql_config")
	log.Println("main")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/site/", siteHandler)
	http.HandleFunc("/admin/", adminHandler)
	log.Println("main2")
	//    http.HandleFunc("/login/", loginHandler)
	http.ListenAndServe(":80", nil)
}
