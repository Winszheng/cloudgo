package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

//路由参数
func main() {
	r := mux.NewRouter()
	//1. 普通路由参数
	r.HandleFunc("/articles/{title}", TitleHandler)

	//2. 正则路由参数，下面例子中限制为英文字母
	r.HandleFunc("/articles/{title:[a-z]+}", TitleHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}

//https://github.com/gorilla/mux#examples
func TitleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // 获取参数
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "title: %v\n", vars["title"])
}
