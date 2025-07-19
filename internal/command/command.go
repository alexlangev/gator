package command

import (
	"errors"
	"fmt"

	s "github.com/alexlangev/gator/internal/state"
)

type Command struct {
	name      string
	arguments []string
}

type Commands struct {
	Commands map[string]func(*s.State, Command) error
}

func (c *Commands) run(s *s.State, cmd Command) error {
	handler, ok := c.Commands[cmd.name]
	if !ok {
		return errors.New("invalid command")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) register(name string, f func(*s.State, Command) error) {
	c.Commands[name] = f
}

func handlerLogin(s *s.State, cmd Command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("login needs an argument")
	}

	// set the user to the given username
	err := s.State.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set!")

	return nil
}
