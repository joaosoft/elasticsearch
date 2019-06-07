package main

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

func queueBulkCreate() {
	// create process manager
	pm := manager.NewManager(manager.WithRunInBackground(false))

	// create queue
	bulkWorkqueueConfig := manager.NewBulkWorkListConfig("queue_001", 100, 10, 2, time.Second*2, manager.FIFO)
	bulkWorkqueue := pm.NewSimpleBulkWorkList(bulkWorkqueueConfig, bulkWorkHandler, bulkWorkRecoverHandler, bulkWorkRecoverWastedRetriesHandler)
	pm.AddWorkList("bulk_queue", bulkWorkqueue)

	if err := bulkWorkqueue.Start(); err != nil {
		log.Errorf("MAIN: error starting bulk workqueue %s", err)
	}

	// add job to queue
	go func() {
		nJobs := 20000
		for i := 1; i <= nJobs; i++ {
			bulkWorkqueue.AddWork(strconv.Itoa(i),
				&person{
					Name: fmt.Sprintf("name %d", i),
					Age:  i,
				})
		}
	}()

	<-time.After(30 * time.Second)
}

func bulkWorkHandler(works []*manager.Work) error {
	log.Infof("handling works with length %d!", len(works))

	bulk := client.Bulk()

	// handle works on elastic bulk
	var err error
	for _, work := range works {
		if err = bulk.Index("persons").Type("person").Id(work.Id).Body(work.Data).DoCreate(); err != nil {
			log.Error(err)
			return err
		}
	}

	result, err := bulk.Execute()
	if err != nil {
		log.Error(err)
		fmt.Printf("%+v", result)
	} else {
		fmt.Printf("success with %+v", result)
	}

	return nil
}

func bulkWorkRecoverHandler(list manager.IList) error {
	fmt.Printf("\nrecovering list with length %d", list.Size())
	return nil
}

func bulkWorkRecoverWastedRetriesHandler(id string, data interface{}) error {
	fmt.Printf("\nrecovering work with id: %s, data: %+v", id, data)
	return nil
}
