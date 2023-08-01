package cron

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("test")
}
func Cron() {
	gocron.Every(10).Second().Do(task)
	//un comment to start running crons
	<-gocron.Start()
}
