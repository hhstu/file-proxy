package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("参数错误： ")
		fmt.Println("    ./file-proxy xxxxx")
		os.Exit(1)
	}
	workPath := os.Args[1]
	http.Handle("/", http.FileServer(http.Dir(workPath)))
	http.ListenAndServe(":8080", nil)
}
