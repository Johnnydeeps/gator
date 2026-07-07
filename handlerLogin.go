package main

import "fmt"

func handlerLogin(appStatePtr *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login requires a username")
	}
	username := cmd.Args[0]

	err := appStatePtr.configPointer.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("user set to: %s\n", username)
	return nil
}
