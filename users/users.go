package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smit2909/secondService/address"
	"github.com/smit2909/secondService/company"
	"github.com/smit2909/secondService/post"
)

type User struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Phone    string          `json:"phone"`
	Addr     address.Address `json:"address"`
	Comp     company.Company `json:"company"`
}

type UserPosts struct {
	User  *User
	Posts *[]post.Post
}

func (up *UserPosts) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["id"] = up.User.Id
	m["userInfo"] = map[string]string{
		"name":     up.User.Name,
		"username": up.User.Username,
		"email":    up.User.Email,
	}

	posts := make([]map[string]interface{}, 0)
	for _, p := range *(up.Posts) {
		temp := map[string]interface{}{
			"id":    p.Id,
			"title": p.Title,
			"body":  p.Body,
		}
		posts = append(posts, temp)
	}

	m["posts"] = posts
	return json.Marshal(m)
}

func RequestUser(userId int, user *User) error {
	// we can make a request with timeout context
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func RequestUserPosts(userId int, posts *[]post.Post) error {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", userId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(posts)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func MergeUserPosts(user *User, posts *[]post.Post) (string, error) {
	up := &UserPosts{
		User:  user,
		Posts: posts,
	}
	data, err := json.Marshal(up)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
