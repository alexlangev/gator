package main

import (
	"fmt"

	"github.com/alexlangev/gator/internal/config"
)

func main() {
	userConfig, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	userConfig.SetUser("Alex")

	userConfig, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userConfig)
}
