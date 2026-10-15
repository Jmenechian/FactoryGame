package main

import (
	"fmt"
	"time"
)

type World struct {
	width  int
	height int
}

func NewWorld(width, height int) *World {
	return &World{width: width, height: height}
}

func (w *World) Step() bool {
	return true
}

type Player struct {
	money int
}

func NewPlayer(money int) Player {
	return Player{money: money}
}

func (p *Player) Status() string {
	return fmt.Sprintf("Player has %d money", p.money)
}

func main() {
	world := NewWorld(10, 10)
	player := NewPlayer(50)
	for i := 0; i < 100; i++ {
        world.Step()
        fmt.Println(player.Status())
		time.Sleep(1 * time.Second)
    }
}