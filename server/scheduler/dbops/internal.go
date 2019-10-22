package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)



func ReadVideoDeletionRecord(count int) ([]string, error) {
	var ids []string
	statement, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	if err != nil {
		return ids, err
	}

	rows, err := statement.Query(count)
	if err != nil {
		log.Printf("Query video deletion record error %v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer statement.Close()
	return ids, nil
}

func DeleteVideoDeleteRecord(vid string) error {
	statement, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id=?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecofd error: %v", err)
		return err
	}
	defer statement.Close()
	return nil
}
