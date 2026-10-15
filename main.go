package main

import (
	"time"
)

func main() {
	game := NewGame()

	game.ExecuteCommand(Command{Name: "help"})

	cmdChan := make(chan Command, 16)
	go ReadCommands(cmdChan)

	for game.running {
		game.ProcessCommands(cmdChan)
		if !game.running {
			break
		}

		game.Step()
		time.Sleep(100 * time.Millisecond)
	}
}