package main

import (
	"log"
	"os"

	"github.com/Johnnydeeps/gator/internal/config"
)

type state struct {
	configPointer *config.Config
}

func main() {
	//**************************Saved Config json to struct in memory logic *********************
	savedConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading saved (on disk) config file: %v", err)
	}
	appState := state{
		configPointer: &savedConfig,
	}
	appStatePtr := &appState
	// ******************************************************************************************

	// builds the list of commands available for use as arguments
	commands := commandsList{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.registerCommand("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("usage: <command> pass an argument")
	}

	//. ********************************command exectuion logic**********************************
	//. os.Args[0] = the program itself (that long /tmp/.../gator path)
	//. os.Args[1] = the first thing the user typed (the command name)
	//. os.Args[2:] = everything after that (the arguments)
	err = commands.run(appStatePtr, command{Name: os.Args[1], Args: os.Args[2:]})
	if err != nil {
		log.Fatal(err)
	}
}
