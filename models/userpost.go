package models

import "fmt"

type UserPost struct {
	Id       int         `json:"id"`
	UserInfo userInfoDto `json:"userInfo"`
	Posts    []postDto   `json:"posts"`
}

func CreateUserPost(user User, posts []Post) *UserPost {
	postsDto := make([]postDto, 0)
	for _, p := range posts {
		fmt.Println(p.Id, p.Title)
		pDto := postDto{Id: p.Id, Title: p.Title, Body: p.Body}
		postsDto = append(postsDto, pDto)
	}

	return &UserPost{
		Id: user.Id,
		UserInfo: userInfoDto{
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
		},
		Posts: postsDto,
	}
}
