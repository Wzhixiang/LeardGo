package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	get()
	post()
	postWithJson()
}

func get() {
	resp, _ := http.Get("http://localhost:8080/login?account=admin&password=123456")
	body, _ := ioutil.ReadAll(resp.Body)

	var result Resp
	//解析json
	json.Unmarshal(body, &result)

	fmt.Println("code = ", result.Code, "msg = ", result.Msg)
}

func post() {
	form := make(url.Values)
	form.Add("account", "admin")
	form.Add("password", "123456")
	resp, _ := http.PostForm("http://127.0.0.1:8080/login", form)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("result = " + string(body))
}

func postWithJson() {
	account := Account{"admin", "123456"}
	jsonAccount, _ := json.Marshal(account)
	resp, _ := http.Post("http://localhost:8080/loginByJson", "application/json", bytes.NewBuffer([]byte(jsonAccount)))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("result = " + string(body))
}

type Account struct {
	Account  string
	Password string
}

type Resp struct {
	Code int
	Msg  string
}
