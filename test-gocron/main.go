package main

import (
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.Local)
	job1, err := s.Every(5).Second().StartAt(time.Now().Add(10 * time.Second)).Do(Jobs1)
	if err != nil {
		panic(err)
	}
	log.Println("Start")
	s.StartAsync()
	time.Sleep(1 * time.Minute)
	s.RemoveByReference(job1)
	_, err = s.Every(3).Second().Do(Jobs2)
	time.Sleep(1 * time.Minute)
	log.Println("End")
}

func Jobs1() {
	log.Println("In job 1")
}

func Jobs2() {
	log.Println("In job 2")
}
