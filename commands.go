package main

import "fmt"

// singular command struct
type command struct {
	Name string
	Args []string
}

// full list of available command structs, which take a pointer to the shared state struct
// in main.go
type commandsList struct {
	registeredCommands map[string]func(*state, command) error
}

// method for commandsList to add a new function to the list
func (c *commandsList) registerCommand(name string, function func(*state, command) error) {
	c.registeredCommands[name] = function
}

// method for retrieving and running a function in the commands list
func (c *commandsList) run(appStatePtr *state, cmd command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found: %s", cmd.Name)
	}
	return handler(appStatePtr, cmd)
}
