package main

import (
	"net/http"

	"github.com/smit2909/secondService/routers"
)

func main() {
	r := routers.SetupRouter()
	http.ListenAndServe(":8000", r)
}
