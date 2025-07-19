package main

import (
	"fmt"
	"os"

	cmd "github.com/alexlangev/gator/internal/command"
	c "github.com/alexlangev/gator/internal/config"
	s "github.com/alexlangev/gator/internal/state"
)

func main() {
	userConfig, err := c.Read()
	if err != nil {
		fmt.Println(err)
	}

	appState := s.State{
		State: &userConfig,
	}

	// userConfig.SetUser()

	// TODO:
	// create an instance of commands struct with an initialized map of handler functions
	// register a handler function for the login command

	inputArgs := os.Args[1:]
	if len(inputArgs) == 0 {
		fmt.Println("Invalid number of arguments")
	}

	switch {
	case inputArgs[0] == "login":
		// call setuser
		// update appState
	default:
		fmt.Println("invalid argument")
		os.Exit(0)
	}

}
