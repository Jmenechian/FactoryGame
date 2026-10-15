package main

import "fmt"

type Player struct {
	money int
	resources map[string]int
}

func NewPlayer(money int) Player {
	return Player{
		money: money,
		resources: make(map[string]int),
	}
}

func (p *Player) AddResource(resource string, amount int) {
	p.resources[resource] += amount
}

func (p *Player) Status() string {
	return fmt.Sprintf("Player has %d money and %d iron", p.money, p.resources["iron"])
}

func (p *Player) Mine(resource string) {
	p.resources[resource]++
}

