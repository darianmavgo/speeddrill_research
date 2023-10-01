package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	fontSize     = 24
)

var (
	guess           string
	randomNumber    int
	guessingStarted bool
)

func update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) && guessingStarted {
		userGuess, err := strconv.Atoi(guess)
		if err == nil {
			if userGuess < randomNumber {
				fmt.Println("Too low! Try again.")
			} else if userGuess > randomNumber {
				fmt.Println("Too high! Try again.")
			} else {
				fmt.Println("Congratulations! You guessed the correct number.")
				guessingStarted = false
			}
		}
	} else {
		for i := ebiten.Key(0); i <= ebiten.KeyMax; i++ {
			if ebiten.IsKeyJustPressed(i) {
				if i >= ebiten.Key0 && i <= ebiten.Key9 {
					guess += string('0' + i - ebiten.Key0)
				} else if i == ebiten.KeyBackspace && len(guess) > 0 {
					guess = guess[:len(guess)-1]
				}
			}
		}
	}

	return nil
}

func draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	text.Draw(screen, guess, mplusFont, 10, 10, color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber = rand.Intn(100) + 1
	guessingStarted = true

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Guess the Number Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
