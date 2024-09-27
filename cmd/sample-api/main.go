package main

import (
	"net/http"

	"github.com/koheiterajima-bs/golang-docker-restapi-sample/controller"
	"github.com/koheiterajima-bs/golang-docker-restapi-sample/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	// HTTPサーバーを定義する構造体
	server := http.Server{
		Addr: ":8080",
	}
	// ルーティング設定(/todos/というパスにアクセスされたときに ro.HandleTodosRequest関数が呼ばれるように設定)
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	// サーバーの起動
	server.ListenAndServe()
}
