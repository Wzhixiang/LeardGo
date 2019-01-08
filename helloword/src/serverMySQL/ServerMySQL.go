package main

import "database/sql"

func main() {
	db, error := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

}
