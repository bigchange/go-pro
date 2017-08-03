package main

import (
	"log"
	"net/http"
)

// 如何导入其他包
import "web/services"

var api = new(services.CaseApi)

func main() {
	http.HandleFunc("/index", api.SayHello)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
