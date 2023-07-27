package cron

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("test", time.DateTime)
}
func Cron() {
	gocron.Every(10).Second().Do(task)
	//un comment to start running crons
	<-gocron.Start()
}
