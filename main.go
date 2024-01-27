package main

import (
	"fmt"

	game "github.com/MrWindwaker/GGJ/Game"
)

func main() {
	fmt.Println("Hello")

	g := game.Get_Game()

	g.Run()
}
