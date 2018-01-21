package schedule

import (
	"log"
	"time"

	"github.com/carlescere/scheduler"
)

func job1() {
	log.Println("job1", time.Now())
}

func Create() error {

	if _, err := scheduler.Every(1).Minutes().Run(job1); err != nil {
		return err
	}

	return nil
}
