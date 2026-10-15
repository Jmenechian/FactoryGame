package main

import "fmt"

type Player struct {
	money int
	iron int
}

func NewPlayer(money int) Player {
	return Player{money: money}
}

func (p *Player) AddResource(resource string, amount int) {
	if resource == "iron" {
		p.iron += amount
	}
}

func (p *Player) Status() string {
	return fmt.Sprintf("Player has %d money and %d iron", p.money, p.iron)
}

func (p *Player) Mine(resource string) {
	if resource == "iron" {
		p.iron++
	}
}

