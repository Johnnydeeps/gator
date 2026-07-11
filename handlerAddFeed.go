package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Johnnydeeps/gator/internal/database"
	"github.com/google/uuid"
)

func addRSSFeedDB(appStatePtr *state, cmd command) error {
	if len(cmd.UserArgs) < 2 {
		return fmt.Errorf("addfeed requires a name and a url")
	}

	user, err := appStatePtr.databasePointer.GetUser(context.Background(), appStatePtr.configPointer.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not in database:%w", err)
	}

	feed, err := appStatePtr.databasePointer.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.UserArgs[0],
		Url:       cmd.UserArgs[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	// auto follow a feed a given user creates
	feedFollow, err := appStatePtr.databasePointer.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s created, %s is now following %s\n", feedFollow.FeedName, feedFollow.UserName, feedFollow.FeedName)
	return nil
}
