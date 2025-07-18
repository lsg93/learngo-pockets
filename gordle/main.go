package main

import (
	"learngo-pockets/gordle/gordle"
	"os"
)

func main() {
	game := gordle.New(os.Stdout, gordle.WithSolution("lucid"), gordle.WithDictionary([]string{
		"audio", "crane", "drive", "steer", "hello", "jello",
		"inter", "climb", "rider", "foxes", "grunt", "mount",
		"zones", "homes", "blaze", "lucid", "bound", "flame",
		"quick", "weird", "alone", "count", "teach", "grind"}))

	game.Play()
}
