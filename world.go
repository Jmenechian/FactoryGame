package main

type World struct {
	width  int
	height int

	miners int
	Player Player
}

func NewWorld(width, height int) *World {
	return &World{
		width:  width,
		height: height,
		Player: NewPlayer(50),
	}
}

func (w *World) Step() bool {
	for range w.miners {
		w.Player.AddResource("iron", 1)
	}
	return true
}
