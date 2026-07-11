package main

import (
	"context"
	"fmt"

	"github.com/Johnnydeeps/gator/internal/database"
)

func handlerUserFollowing(appStatePtr *state, cmd command, user database.User) error {
	if len(cmd.UserArgs) != 0 {
		return fmt.Errorf("following usage: following ")
	}

	follows, err := appStatePtr.databasePointer.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}
	return nil
}
