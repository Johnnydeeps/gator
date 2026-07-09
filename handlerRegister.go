package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Johnnydeeps/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(appStatePtr *state, cmd command) error {
	if len(cmd.UserArgs) == 0 {
		return fmt.Errorf("register requires a username as an argument")
	}
	username := cmd.UserArgs[0]

	user, err := appStatePtr.databasePointer.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		fmt.Printf("error creating user: %s", err)
		os.Exit(1)
	}

	err = appStatePtr.configPointer.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("user: %s registered successfully\n", username)
	log.Printf("%+v", user)
	return nil
}
