package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/0g3/treasure-app/app/config"
	"os"
	"os/exec"
)

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
	case "create-db":
		execute("mysql",
			"-u", config.Env().DBUser,
			"-h", config.Env().DBAddress,
			"-p"+config.Env().DBPassword,
			"--protocol", "tcp",
			"-e", "create database "+config.Env().DBName,
		)

	case "drop-db":
		execute("mysql",
			"-u", config.Env().DBUser,
			"-h", config.Env().DBAddress,
			"-p"+config.Env().DBPassword,
			"--protocol", "tcp",
			"-e", "drop database "+config.Env().DBName,
		)

	case "get-cli-cmd":
		fmt.Printf(
			"mysql -u %s -h %s -p%s --protocol tcp %s",
			config.Env().DBUser,
			config.Env().DBAddress,
			config.Env().DBPassword,
			config.Env().DBName,
		)
	}
}
