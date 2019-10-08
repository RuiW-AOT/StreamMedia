package dbops

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/RuiW-AOT/StreamMedia/server/api/defs"
	"github.com/RuiW-AOT/StreamMedia/server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	statement, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		log.Printf("Create user error %s", err)
		return err
	}

	_, err = statement.Exec(loginName, pwd)
	if err != nil {
		log.Printf("failed to create user: %s", err)
		return errors.Wrap(err, "failed to create user")
	}
	defer statement.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	statement, err := dbConn.Prepare("Select pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("Get user error %s", err)
		return "", err
	}
	var pwd string
	err = statement.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Cannot get the user: %s", err)
		return "", errors.Wrap(err, "failed to get user")
	}
	defer statement.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	statement, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? and pwd=?")
	if err != nil {
		log.Printf("Delete user error %s", err)
		return err
	}
	_, err = statement.Exec(loginName, pwd)
	if err != nil {
		log.Printf("failed to delete user: %s", err)
		return errors.Wrap(err, "failed to delete user")
	}
	defer statement.Close()
	return nil
}

func AddNewVideo(authorID int, name string) (*defs.VideoInfo, error) {
	// create Uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	statement, err := dbConn.Prepare(`INSERT INTO videos
	 (id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(vid, authorID, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{ID: vid, AuthorID: authorID, Name: name, DisplayCtime: ctime}
	defer statement.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	statement, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM videos WHERE id=?")

	var aid int
	var dct string
	var name string

	err = statement.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer statement.Close()

	res := &defs.VideoInfo{ID: vid, AuthorID: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	statement, err := dbConn.Prepare("DELETE FROM videos WHERE id=?")
	if err != nil {
		return err
	}

	_, err = statement.Exec(vid)
	if err != nil {
		return err
	}

	defer statement.Close()
	return nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	statement, err := dbConn.Prepare(`SELECT videos.id, videos.author_id, videos.name, videos.display_ctime FROM videos
		INNER JOIN users ON videos.author_id = users.id
		WHERE users.login_name = ? AND videos.create_time > FROM_UNIXTIME(?) AND videos.create_time <= FROM_UNIXTIME(?) 
		ORDER BY videos.create_time DESC`)

	var res []*defs.VideoInfo

	if err != nil {
		return res, err
	}

	rows, err := statement.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{ID: id, AuthorID: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}

	defer statement.Close()

	return res, nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	statement, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer statement.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
		ORDER BY comments.time DESC`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{ID: id, VideoID: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}
