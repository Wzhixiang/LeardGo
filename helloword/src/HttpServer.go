package main

import (
	"log"
	"net/http"
)

func helloServer(write http.ResponseWriter, request *http.Request) {
	//io.WriteString(write, "hello world!")

	write.Write([]byte("hello!"))
}

func main() {
	http.HandleFunc("/hello", helloServer)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
