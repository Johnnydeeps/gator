package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(appStatePtr *state) error {
	nextFetch, err := appStatePtr.databasePointer.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = appStatePtr.databasePointer.MarkFeedFetched(context.Background(), nextFetch.ID)
	if err != nil {
		return err
	}

	fetchedFeed, err := fetchRSSFeed(context.Background(), nextFetch.Url)
	if err != nil {
		return err
	}

	for _, feed := range fetchedFeed.Channel.Item {
		fmt.Printf("Feed Title: %s\n", feed.Title)
	}
	return nil
}
