/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"

	"github.com/seungjinyu/gitcode/cmd"
)

func main() {
	os.Setenv("VERSION", "1.0.0")
	cmd.Execute()
}
