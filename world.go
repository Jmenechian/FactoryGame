package main

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
