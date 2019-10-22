package tasksrunner

import (
	"errors"
	"log"

	"github.com/RuiW-AOT/StreamMedia/server/scheduler/dbops"
)

func deleteVideo(vid string) error {
	err := os.Remove("./videos/"+vid)
	if err != nil  && !os.IsNotExist(err){
		return err
	}
	return nil
}


func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video cear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecuter(dc dataChan) error {
	errmap := &sync.map{}
	var err error
	forloop:
		for {
			select {
			case vid := <- dc:
				go func(id interfacr{}) {
					if err := deleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err :=dbops.DeleteVideoDeleteRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
				default:
					break forloop
			}
		}
	errMap.Range(func(k, v interface{})bool{
		err = v.(error)
		if err != nil {
			return false
		}
	})

	return err

}
