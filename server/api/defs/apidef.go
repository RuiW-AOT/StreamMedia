package defs

type UserCredentials struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
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
