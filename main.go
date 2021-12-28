package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smit2909/secondService/post"
	"github.com/smit2909/secondService/users"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("hello world")
	errGrp, _ := errgroup.WithContext(context.Background())
	user := &users.User{}
	posts := make([]post.Post, 0)
	errGrp.Go(
		func() error {
			return users.RequestUser(1, user)
		},
	)
	errGrp.Go(
		func() error {
			return users.RequestUserPosts(1, &posts)
		},
	)
	err := errGrp.Wait()
	if err != nil {
		log.Fatal(err)
		return
	}

	finalOp, err := users.MergeUserPosts(user, &posts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(finalOp)
}
