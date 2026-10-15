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
	fmt.Println("Available commands: status, mine [resource], sell, hire, buy, quit, help")
	return nil
}

func handleStatus(g *Game, args []string) error {
	fmt.Printf("[Tick %d] %s (Miners: %d)\n", g.tick, g.world.Player.Status(), g.world.miners)
	return nil
}

func handleBuy(g *Game, args []string) error {
	g.world.Player.money += 10
	fmt.Printf("Added money! %s\n", g.world.Player.Status())
	return nil
}

func handleQuit(g *Game, args []string) error {
	g.running = false
	fmt.Println("Exiting simulation...")
	return nil
}

func handleMine(g *Game, args []string) error {
	var resource string
	if len(args) > 0 {
		resource = args[0]
	}

	if resource == "" {
		return fmt.Errorf("no resource specified")
	}

	g.world.Player.Mine(resource)
	fmt.Printf("Mined %s! %s\n", resource, g.world.Player.Status())
	return nil
}

func handleSell(g *Game, args []string) error {
	var resource string
	if len(args) > 0 {
		resource = args[0]
	}

	if resource == "" {
		return fmt.Errorf("no resource specified")
	}

	ammount_to_sell := g.world.Player.resources[resource]
	g.world.Player.resources[resource] -= ammount_to_sell
	g.world.Player.money += ammount_to_sell * 5
	fmt.Printf("Sold %d %s! %s\n", ammount_to_sell, resource, g.world.Player.Status())

	return nil
}

func handleHire(g *Game, args []string) error {
	if g.world.Player.money >= 50 {
		g.world.miners++
		g.world.Player.money -= 50
		fmt.Printf("Hired a miner! Total miners: %d (%s)\n", g.world.miners, g.world.Player.Status())
	} else {
		fmt.Println("Not enough money to hire a miner (costs 50).")
	}
	return nil
}
