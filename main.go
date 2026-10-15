package main

import (
	"time"
)

func main() {
	game := NewGame()

	cmdChan := make(chan Command, 16)
	go ReadCommands(cmdChan)

	for game.running && game.tick < 100 {
		game.ProcessCommands(cmdChan)
		if !game.running {
			break
		}

		game.Step()
		time.Sleep(1 * time.Second)
	}
}