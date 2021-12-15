/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/seungjinyu/gitcode/cmd"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env file import error")
	}
	cmd.Execute()
}
