package view

type Meme struct {
	Title             string    `json:"title"`
	Id                int64     `json:"id"`
	Tag               *[]string `json:"tag"`
	Username          string    `json:"username"`
	UsernameAvatarUrl string    `json:"username_avatar_url"`
	Picture           string    `json:"picture"`
	Like              int       `json:"like"`
}
