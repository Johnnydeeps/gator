package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Johnnydeeps/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(appStatePtr *state, cmd command) error {
	if len(cmd.UserArgs) != 1 {
		return fmt.Errorf("follow usage: follow <url>")
	}
	urlToFollow := cmd.UserArgs[0]

	userconfig := appStatePtr.configPointer.CurrentUserName

	userDB, err := appStatePtr.databasePointer.GetUser(context.Background(), userconfig)
	if err != nil {
		return err
	}

	feedDB, err := appStatePtr.databasePointer.GetFeedByURL(context.Background(), urlToFollow)
	if err != nil {
		return err
	}

	feedFollow, err := appStatePtr.databasePointer.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userDB.ID,
		FeedID:    feedDB.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("user: %s now following feed: %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
