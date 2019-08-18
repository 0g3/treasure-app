package main

import (
	"github.com/0g3/treasure-app/app/server"
	"log"
)

func main() {
	srv, err := server.New()
	if err != nil {
		log.Fatal("サーバーの生成に失敗:", err)
	}
	log.Fatal("サーバーの起動に失敗:", srv.Run())
}
