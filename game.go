package main

import "fmt"

type Game struct {
	world    *World
	running  bool
	tick     int
	handlers map[string]CommandHandler
}

func NewGame() *Game {
	g := &Game{
		world:    NewWorld(10, 10),
		running:  true,
		handlers: make(map[string]CommandHandler),
	}
	g.registerHandlers()
	return g
}

func (g *Game) registerHandlers() {
	g.handlers["help"] = handleHelp
	g.handlers["status"] = handleStatus
	g.handlers["buy"] = handleBuy
	g.handlers["exit"] = handleQuit
	g.handlers["quit"] = handleQuit
	g.handlers["mine"] = handleMine
	g.handlers["sell"] = handleSell
	g.handlers["hire"] = handleHire
}

func (g *Game) ExecuteCommand(cmd Command) {
	handler, exists := g.handlers[cmd.Name]
	if !exists {
		fmt.Printf("Unknown command: %q. Type 'help' for available commands.\n", cmd.Name)
		return
	}
	if err := handler(g, cmd.Args); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (g *Game) ProcessCommands(commands <-chan Command) {
	for {
		select {
		case cmd, ok := <-commands:
			if !ok {
				g.running = false
				return
			}
			g.ExecuteCommand(cmd)
		default:
			return
		}
	}
}

func (g *Game) Step() {
	g.world.Step()
	g.tick++
}
