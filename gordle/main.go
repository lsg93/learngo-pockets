package main

import (
	"learngo-pockets/gordle/gordle"
	"os"
)

func main() {
	game := gordle.New(os.Stdout, gordle.WithSolution("lucid"))
	game.Play()
}
