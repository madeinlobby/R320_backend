package view

import "time"

type Meme struct {
	Title    string    `json:"title"`
	Id       int64     `json:"id"`
	Tag      []string  `json:"tag"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
	Picture  string    `json:"picture"`
	Like     int       `json:"like"`
	Date     time.Time `json:"date"`
}
type Comment struct {
	ID      int64      `json:"id"`
	Author  string     `json:"author"`
	Avatar  string     `json:"avatar"`
	Text    string     `json:"body"`
	Like    int        `json:"like"`
	Replies *[]Comment `json:"replies"`
	Date    time.Time  `json:"date"`
}
