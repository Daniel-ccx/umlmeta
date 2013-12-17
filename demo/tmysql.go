package main

import (
	"database/sql"
	"fmt"
	_ "mysql"
	"strconv"
)

func main() {
	db, err := sql.Open("mysql", "root:123qwe@tcp(localhost:3306)/wordpress?charset=utf8")
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("select `id`, user_nicename from wp_users")
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
    id := 0
    user_nickname := "" 
	for rows.Next() {
		rerr := rows.Scan(&id, &user_nickname)
		if rerr == nil {
			fmt.Println(strconv.Itoa(id), user_nickname)
		}
	}

	//db.Close()
}
