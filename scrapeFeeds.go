package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Johnnydeeps/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

	for _, item := range fetchedFeed.Channel.Item {
		// helper function from the time standard libray to convert whatever input string is returned from
		// the RSS feed to Go's standard time.Time format to remove any ambiguity of how time is kept.
		// time.RFC1123Z is go's layout for RSS feeds, item.pubdate is the string in the feed.item[] slice
		// return from the http request.
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println("couldn't parse published date:", err)
			continue
		}
		//**************************************************************************************************
		_, err = appStatePtr.databasePointer.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			// The description column from the 005_posts.sql schema allows for a NULL (no value at all),
			// which is different from an empty string (a value that just happens to be blank).
			// sql.NullString is the Go type that lets you explicitly say which one you mean each time
			// you insert a row.
			// If the RSS item had a real description, item.Description != "" is true → Valid: true →
			// the database will store the actual text.
			// If the RSS item had no description (item.Description is ""),
			// then item.Description != "" is false → Valid: false → the database will store NULL instead
			// of an empty string.
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			// ******************************************************************************************
			PublishedAt: publishedAt,
			FeedID:      nextFetch.ID,
		})
		if err != nil {
			// errors.As is a function in Go's standard library, this is comparing any errors that happen
			// in the pq database driver (go postgres driver) with any err that have been caught in the
			// post creation above. in this case, this would be the error "23505" which is the database
			// error that occurs when a unique value in this case a url is already present in the database.
			// So this compares go's captured err, with the *pq.Error and if that code == "23505" to continue
			// without raising or logging the error.
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				continue
			}
			log.Println("couldn't create post:", err)
			continue
		}
	}
	return nil
}
