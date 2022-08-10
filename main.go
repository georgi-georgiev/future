package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.TagsUnique()

	_, _ = s.Every(1).Week().Tag("foo").Do(task)
	_, err := s.Every(1).Week().Tag("foo").Do(task)
	if err != nil {
		panic(err)
	}

	s2 := gocron.NewScheduler(time.UTC)

	s2.Every(2).Day().Tag("tag").At("10:00").Do(task)
	s2.Every(1).Minute().Tag("tag").Do(task)
	s2.RunByTag("tag")
	// both jobs will run
}

func task() {
	fmt.Println("task")
}
