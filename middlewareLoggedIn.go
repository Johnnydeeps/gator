package main

import (
	"context"

	"github.com/Johnnydeeps/gator/internal/database"
)

func middlewareLoggedIn(handler func(appStatePtr *state, cmd command, user database.User) error) func(*state, command) error {
	return func(appStatePtr *state, cmd command) error {
		user, err := appStatePtr.databasePointer.GetUser(context.Background(),
			appStatePtr.configPointer.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(appStatePtr, cmd, user)
	}
}
