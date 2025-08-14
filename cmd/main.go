package main

import (
	customhttp "backend_service/internal/http"
	"fmt"
	"net/http"
)

func main() {
	r := customhttp.NewRouter()

	fmt.Println("Backend service running on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}
}
