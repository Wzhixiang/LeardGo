package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginByJson", loginByJson)
	http.ListenAndServe(":8080", nil)
}

/*
	get|post
	接收x-www-form-urlencoded类型请求
	通过request.Form["参数名"]获取参数值
*/
func login(writer http.ResponseWriter, request *http.Request) {
	//解析参数，默认是不会解析
	request.ParseForm()

	//todo account数组， aError是否有传参数
	account, aError := request.Form["account"]
	password, pError := request.Form["password"]

	var result = Resp{}

	if aError && pError {
		//
		if account[0] == "admin" && password[0] == "123456" {
			result.Code = 200
			result.Msg = "登录成功"
		} else {
			result.Code = 201
			result.Msg = "账号或密码错误"
		}
	} else {
		result.Code = 201
		result.Msg = "账号或密码错误"
	}

	json.NewEncoder(writer).Encode(result)
}

/*
	post
	接收application|json类型请求
	需要在Request.Body中获取json
*/
func loginByJson(writer http.ResponseWriter, request *http.Request) {
	var account Account
	//TODO 从request.Body中获取json，然后解析赋值给变量account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		request.Body.Close()
	}
	var result = Resp{}
	if account.Account == "admin" && account.Password == "123456" {
		result.Code = 200
		result.Msg = "登录成功"
	} else {
		result.Code = 201
		result.Msg = "账号或密码错误"
	}

	json.NewEncoder(writer).Encode(result)
}

type Account struct {
	Account  string
	Password string
}

type Resp struct {
	Code int
	Msg  string
}
