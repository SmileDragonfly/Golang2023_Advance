package main

import (
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.Local)
	job, err := s.Every(1).Day().At("22:07;22:09").Do(Jobs)
	if err != nil {
		panic(err)
	}
	log.Println(job.IsRunning())
	s.StartBlocking()
}

func Jobs() {
	log.Println("In job")
}
