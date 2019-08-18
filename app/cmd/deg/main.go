package main

import (
	"github.com/0g3/treasure-app/app/config"
	"log"
	"os"
)

func main() {
	f, err := os.Create(".env")
	if err != nil {
		log.Fatal("ファイル生成に失敗:", err)
	}
	defer f.Close()

	for _, name := range config.Env().GetPrefixedFieldNames() {
		_, err := f.Write([]byte(name + "=\n"))
		if err != nil {
			log.Fatal("書き込みエラー:", err)
		}
	}
}
