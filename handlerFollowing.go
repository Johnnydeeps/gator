package main

import (
	"context"
	"fmt"
)

func handlerUserFollowing(appStatePtr *state, cmd command) error {
	if len(cmd.UserArgs) != 0 {
		return fmt.Errorf("following usage: following ")
	}

	userconfig := appStatePtr.configPointer.CurrentUserName

	userDB, err := appStatePtr.databasePointer.GetUser(context.Background(), userconfig)
	if err != nil {
		return err
	}

	follows, err := appStatePtr.databasePointer.GetFeedFollowsForUser(context.Background(), userDB.ID)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}
	return nil
}
