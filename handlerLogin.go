package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(appStatePtr *state, cmd command) error {
	if len(cmd.UserArgs) == 0 {
		return fmt.Errorf("login requires a username")
	}
	username := cmd.UserArgs[0]

	_, err := appStatePtr.databasePointer.GetUser(context.Background(), username)
	if err != nil {
		fmt.Println("error: user does not exist")
		os.Exit(1)
	}

	err = appStatePtr.configPointer.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("user set to: %s\n", username)
	return nil
}
