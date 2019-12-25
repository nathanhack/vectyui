package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	if _, err := os.Stat("./static/wasm_exec.js"); os.IsNotExist(err) {
		panic("File not found ./static/wasm_exec.js : find it in $GOROOT/misc/wasm/ note it must be from the same version of go used during compiling")
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}