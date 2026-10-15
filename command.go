package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func ParseCommand(raw string) (Command, error) {
	fields := strings.Fields(raw)
	if len(fields) == 0 {
		return Command{}, fmt.Errorf("empty command")
	}
	return Command{
		Name: strings.ToLower(fields[0]),
		Args: fields[1:],
	}, nil
}

func ReadCommands(commands chan<- Command) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd, err := ParseCommand(scanner.Text())
		if err != nil {
			continue
		}
		commands <- cmd
	}
	close(commands)
}

type CommandHandler func(g *Game, args []string) error

func handleHelp(g *Game, args []string) error {
	fmt.Println("Available commands: status, buy, quit, help")
	return nil
}

func handleStatus(g *Game, args []string) error {
	fmt.Printf("[Tick %d] %s\n", g.tick, g.player.Status())
	return nil
}

func handleBuy(g *Game, args []string) error {
	g.player.money += 10
	fmt.Printf("Added money! %s\n", g.player.Status())
	return nil
}

func handleQuit(g *Game, args []string) error {
	g.running = false
	fmt.Println("Exiting simulation...")
	return nil
}
