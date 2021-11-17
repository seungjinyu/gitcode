package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fp, err := os.Open("./.git/config")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	// buffer := make([]byte, 10)

	// r := bufio.NewReader(fp)
	// b, err := r.Peek(1000)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("5 bytes: %s\n", string(b))

	// Splits on newlines by default.
	res, err := findString(fp, "url")

	if err != nil {
		fmt.Println(err)
	}

	giturl := splitString(res)
	fmt.Println(giturl)

	openBrowser(giturl)

}
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()

	default:
		err = fmt.Errorf("unsuppored platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func splitString(res string) string {
	slice := strings.Split(res, " ")
	return slice[2]
}

func findString(fp *os.File, str string) (string, error) {
	scanner := bufio.NewScanner(fp)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), str) {
			resultStr := scanner.Text()
			return resultStr, nil
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
		fmt.Println(err)
	}

	return "Nothing", nil
}
