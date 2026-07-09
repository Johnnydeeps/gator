package main

import (
	"context"
	"fmt"
	"os"
)

func handlerListFeeds(appStatePtr *state, cmd command) error {
	if len(cmd.UserArgs) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := appStatePtr.databasePointer.GetAllFeeds(context.Background())
	if err != nil {
		fmt.Printf("error retrieving feeds from database: %s", err)
		os.Exit(1)
	}

	for _, feed := range feeds {
		user, err := appStatePtr.databasePointer.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("feed name: %s\n", feed.Name)
		fmt.Printf("feed url: %s\n", feed.Url)
		fmt.Printf("user name: %s\n", user.Name)
	}

	return nil
}
