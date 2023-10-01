package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"fmt"
	"log"
)

type word struct {
	text string
	y    float64
}

type game struct {
	words []word
	lives int
}

func (g *game) update() {
	// Move the words down the screen.
	for i := range g.words {
		g.words[i].y += 1
	}

	// Check if any words have hit the bottom of the screen.
	for i := range g.words {
		if g.words[i].y > ebiten.ScreenHeight() {
			// Remove the word from the game.
			g.words = append(g.words[:i], g.words[i+1:]...)

			// Lose a life.
			g.lives--

			// Check if the game is over.
			if g.lives <= 0 {
				// Game over!
				fmt.Println("Game over!")
				ebiten.Quit()
			}
		}
	}
}

func (g *game) draw(screen *ebiten.Image) {
	// Draw the words on the screen.
	for i := range g.words {
		screen.DrawText(g.words[i].text, ebiten.FontDefault, int(g.words[i].y), 10, ebiten.Black)
	}
}

func main() {
	// Create a new Ebiten game.
	game, err := ebiten.NewGame()
	if err != nil {
		// Handle error.
		log.Fatal(err)
	}

	// Create a slice of sample words.
	sampleWords := []string{"hello", "world", "this", "is", "a", "test"}

	// Initialize the game.
	g := game{
		words: make([]word, 0),
		lives: 5,
	}

	// Start the game loop.
	for {
		// Update the game.
		g.update()

		// Draw the game.
		g.draw(game.Screen())

		// Check if the game should quit.
		if ebiten.IsDrawingFinished() {
			break
		}

		// Check if a key is pressed.
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			// The player pressed the A key.
			// Check if there is a word at the top of the screen that starts with the letter A.
			if len(g.words) > 0 && g.words[0].text[0] == 'A' {
				// Remove the word from the game.
				g.words = append(g.words[:0], g.words[1:]...)
			}
		}

		// ... Check for other keys ...
	}
}
