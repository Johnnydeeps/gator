package main

import (
	"context"
	"fmt"

	"github.com/Johnnydeeps/gator/internal/database"
)

func unfollowFeed(appStatePtr *state, cmd command, user database.User) error {
	if len(cmd.UserArgs) != 1 {
		return fmt.Errorf("unfollow usage: unfollow <url>")
	}

	feedID, err := appStatePtr.databasePointer.GetFeedByURL(context.Background(), cmd.UserArgs[0])
	if err != nil {
		return err
	}
	err = appStatePtr.databasePointer.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		FeedID: feedID.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
