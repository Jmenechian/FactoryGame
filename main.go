package main

import (
	// "time"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

// func main() {
// 	game := NewGame()

// 	game.ExecuteCommand(Command{Name: "help"})

// 	cmdChan := make(chan Command, 16)
// 	go ReadCommands(cmdChan)

// 	for game.running {
// 		game.ProcessCommands(cmdChan)
// 		if !game.running {
// 			break
// 		}

// 		game.Step()
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
}