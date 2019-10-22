package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddVideoDeletionRecord(vid string) error {
	statement, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer statement.Close()
	return nil
}
