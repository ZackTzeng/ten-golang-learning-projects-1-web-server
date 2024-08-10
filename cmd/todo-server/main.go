package main

import (
	"fmt"
	"net/http"

	"github.com/ZackTzeng/ten-golang-learning-projects-1-web-server/api/router"
)

func main() {
	port := ":3000"
	r := router.NewRouter()
	fmt.Println("Todo API is served on port", port)
	http.ListenAndServe(port, r)
}
