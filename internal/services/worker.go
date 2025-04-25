package services

import (
	"log"
)

func StartWorkers(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			log.Printf("[Worker-%d] started\n", id)
			for entry := range LogJobs {
				log.Printf("[Worker-%d] processing log from %s", id, entry.Service)

				DB.Create(&entry)
				SendToSlack(entry)
				SendToElasticsearch(entry)
				SendMail(entry)
			}
		}(i + 1)
	}
}
