package routers

import (
	"github.com/gorilla/mux"
	"github.com/smit2909/secondService/controllers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/user-posts/{id}", controllers.UserPostsHandler)
	return r
}
