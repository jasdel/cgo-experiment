package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Get("https://example.com")
	fmt.Println("vim-go")
}
