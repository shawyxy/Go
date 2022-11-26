package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	testClientPost()
}
func testClientPost() {
	//创建客户端
	client := http.Client{}
	//构建参数
	contentType := "application/json;charset=utf-8"

	type LoginInfo struct {
		Username string `json:"username"`
		Password string `json:"Password"`
	}
	login := LoginInfo{Username: "zhb", Password: "123"}
	info, err := json.Marshal(login)
	if err != nil {
		log.Println("json format error:", err)
		return
	}

	body := bytes.NewBuffer(info)
	response, err := client.Post("http://127.0.0.1:9099/login", contentType, body)

	CheckErr32(err)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		fmt.Println("网络请求成功")

		//处理
		data, err := io.ReadAll(response.Body)
		CheckErr32(err)
		fmt.Println(string(data))
	}
}

// 检查错误
func CheckErr32(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常：", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}

//package main
//
//import (
//"encoding/json"
//"fmt"
//"log"
//"net/http"
//"time"
//)
//
//func main() {
//	http.HandleFunc("/login", jsonReqRsp)
//	s := &http.Server{
//		Addr:           ":9090",
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//	err := s.ListenAndServe()
//	if err != nil {
//		log.Fatal("listenAndServe: ", err)
//	} else {
//		fmt.Println("Server listening on 9099！")
//	}
//
//}
//
//type ResData struct {
//	Code     int    `json:"code"`
//	Message  string `json:"message"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//type LoginInfo struct {
//	Username string `json:"username"`
//	Password string `json:"Password"`
//}
//
//// curl  -H "Content-Type:application/json" -X POST  -d  '{"opttype":1 ,"data": { "systemName": "systemName111", "account": ["zgx001", "ZGX00000", "ZGX00001"]} }' -i  'http://127.0.0.1:9090/x'
//func jsonReqRsp(w http.ResponseWriter, r *http.Request) {
//	var reqdata LoginInfo
//
//	// 调用json包的解析，解析请求body
//	if err := json.NewDecoder(r.Body).Decode(&reqdata); err != nil {
//		r.Body.Close()
//		log.Fatal(err)
//	}
//	res := ResData{Code: 404,
//		Message:  "username or password error",
//		Username: reqdata.Username,
//		Password: reqdata.Password}
//	if reqdata.Username == "zhb" && reqdata.Password == "123" {
//		jsonStr, _ := json.Marshal(reqdata)
//		log.Println("req json: ", string(jsonStr))
//
//		res = ResData{Code: 200, Message: "login success", Username: "zhb", Password: "123"}
//	}
//
//	// 返回json字符串给客户端
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(res)
//}
