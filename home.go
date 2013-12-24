package main

import (
	"log"
	"net/http"
)

func main() {
	//var cfg = models.ParseCfg("mysql_config")
	log.Println("main")
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/site/", siteHandler)
	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/login/", loginHandler)
	log.Println("main2")

	http.ListenAndServe(":80", nil)
}
