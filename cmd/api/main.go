package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("LMS API Server is running...")
	http.ListenAndServe(":8080", nil)
}
