package apiHelpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smit2909/secondService/models"
)

func FetchUserInfo(userId int, user *models.User) error {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("could not find resource")
	}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return err
	}
	return nil
}

func FetchUserPosts(userId int, posts *[]models.Post) error {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", userId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("could not find resource")
	}
	err = json.NewDecoder(resp.Body).Decode(posts)
	if err != nil {
		return err
	}
	return nil
}
