package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Johnnydeeps/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(appStatePtr *state, cmd command, user database.User) error {
	if len(cmd.UserArgs) != 1 {
		return fmt.Errorf("follow usage: follow <url>")
	}
	urlToFollow := cmd.UserArgs[0]

	feedDB, err := appStatePtr.databasePointer.GetFeedByURL(context.Background(), urlToFollow)
	if err != nil {
		return err
	}

	feedFollow, err := appStatePtr.databasePointer.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedDB.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("user: %s now following feed: %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
