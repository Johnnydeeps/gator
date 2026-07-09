package main

import (
	"context"
	"fmt"
	"os"
)

func handlerGetAllUsers(appStatePtr *state, cmd command) error {

	users, err := appStatePtr.databasePointer.GetAllUsers(context.Background())
	if err != nil {
		fmt.Printf("error retrieving users from database: %s", err)
		os.Exit(1)
	}

	for _, user := range users {
		if user.Name == appStatePtr.configPointer.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
