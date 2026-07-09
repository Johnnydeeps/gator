package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Johnnydeeps/gator/internal/config"
	"github.com/Johnnydeeps/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	databasePointer *database.Queries
	configPointer   *config.Config
}

func main() {
	//************ Saved Config json to struct in memory logic & DATABASE CONNECTION *************
	//opens the saved on disk .gatorconfig.json and stores in memory as savedConfig
	savedConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading saved (on disk) config file: %v", err)
	}
	// database connection creates a connection in memory with a pointer
	db, err := sql.Open("postgres", savedConfig.DBURL)
	if err != nil {
		log.Fatalf("error reading database location saved (on disk) config file: %v", err)
	}
	// This takes that raw connection handle (db) and wraps it inside your
	// generated database.Queries type — the one SQLC created
	dbQueries := database.New(db)

	// state struct containing the config of the program in memory to be passed into functions
	appState := state{
		configPointer:   &savedConfig,
		databasePointer: dbQueries,
	}
	appStatePtr := &appState
	// *******************************************************************************************

	// builds the list of commands available for use as arguments
	commands := commandsList{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.registerCommand("login", handlerLogin)
	commands.registerCommand("register", handlerRegister)
	commands.registerCommand("reset", handlerResetDB)

	if len(os.Args) < 2 {
		log.Fatal("usage: <command> pass an argument")
	}

	//. ********************************command exectuion logic**********************************
	//. os.Args[0] = the program itself (that long /tmp/.../gator path)
	//. os.Args[1] = the first thing the user typed (the command name)
	//. os.Args[2:] = everything after that (the arguments)
	err = commands.run(appStatePtr, command{Name: os.Args[1], UserArgs: os.Args[2:]})
	if err != nil {
		log.Fatal(err)
	}
}
