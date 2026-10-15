package main

import "fmt"

type Player struct {
	money int
}

func NewPlayer(money int) Player {
	return Player{money: money}
}

func (p *Player) Status() string {
	return fmt.Sprintf("Player has %d money", p.money)
}
