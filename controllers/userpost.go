package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/smit2909/secondService/apiHelpers"
	"github.com/smit2909/secondService/models"
	"golang.org/x/sync/errgroup"
)

func UserPostsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "not a valid user id", http.StatusBadRequest)
		return
	}
	user := &models.User{}
	posts := make([]models.Post, 0)

	errGrp, _ := errgroup.WithContext(r.Context())
	errGrp.Go(
		func() error {
			return apiHelpers.FetchUserInfo(id, user)
		},
	)
	errGrp.Go(
		func() error {
			return apiHelpers.FetchUserPosts(id, &posts)
		},
	)
	err = errGrp.Wait()

	if err != nil {
		http.Error(w, "some error occurred", http.StatusInternalServerError)
		return
	}

	userPost := models.CreateUserPost(*user, posts)
	err = json.NewEncoder(w).Encode(userPost)
	if err != nil {
		http.Error(w, "unable to serialize the data", http.StatusInternalServerError)
		return
	}
}
