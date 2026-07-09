package main

import (
	"context"
	"fmt"
	"os"
)

func handlerResetDB(appStatePtr *state, cmd command) error {

	err := appStatePtr.databasePointer.ResetDB(context.Background())
	if err != nil {
		fmt.Printf("error reseting user database: %s", err)
		os.Exit(1)
	}
	fmt.Println("users database successful reset")
	return nil
}
