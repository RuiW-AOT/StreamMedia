package defs

type UserCredentials struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	ID           string `json:"id"`
	AuthorID     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	ID      string `json:"id"`
	VideoID string `json:"video_id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type SimpleSession struct {
	Username string
	TTL      int64
}

type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type UserInfo struct {
	Id int `json:"id"`
}

type NewComment struct {
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
}

type NewVideo struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}
