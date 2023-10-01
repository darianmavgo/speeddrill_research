package main

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"

	"fmt"
)

func main() {

	// Change the TPS to 30.
	ebiten.SetTPS(30)

	// Print the new TPS.
	fmt.Println(ebiten.TPS())

	// Start the game.
	// game.Run()
}
