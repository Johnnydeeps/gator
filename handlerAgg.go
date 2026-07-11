package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func handlerAgg(appStatePtr *state, cmd command) error {
	if len(cmd.UserArgs) != 1 {
		fmt.Println("usage: <time_between_reqs>")
		os.Exit(1)
		return nil
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.UserArgs[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		err := scrapeFeeds(appStatePtr)
		if err != nil {
			log.Println(err)
		}
	}
}
