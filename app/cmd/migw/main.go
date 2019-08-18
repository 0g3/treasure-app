package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/0g3/treasure-app/app/config"
	"os"
	"os/exec"
	"path"
)

const (
	createUsage = "new example: go run cmd/migw/main.go new create_users"
	upUsage     = "up example: go run cmd/migw/main.go up"
	downUsage   = "down example: go run cmd/migw/main.go down"
)

var (
	dbFmt = fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env().DBUser,
		config.Env().DBPassword,
		config.Env().DBAddress,
		config.Env().DBPort,
		config.Env().DBName,
	)
	fileFmt string
)

func init() {
	current, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileFmt = "file://" + path.Join(current, config.Env().MigrationDir)
}

func execute(name string, arg ...string) {
	cmd := exec.Command(
		name,
		arg...,
	)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("err:", err)
		fmt.Println(stderr.String())
		os.Exit(1)
	}
	fmt.Println(stdout.String())
}

func main() {
	flag.Parse()
	args := flag.Args()
	switch args[0] {
	case "new":
		if len(args) != 2 {
			fmt.Println(createUsage)
			os.Exit(1)
		}
		execute("migrate", "create", "-dir", config.Env().MigrationDir, "-ext", "sql", args[1])

	case "up":
		execute("migrate", "-source", fileFmt, "-database", dbFmt, "up")

	case "down":
		execute("migrate", "-source", fileFmt, "-database", dbFmt, "down")

	default:
		fmt.Println(createUsage)
		fmt.Println(upUsage)
		fmt.Println(downUsage)
	}
}
