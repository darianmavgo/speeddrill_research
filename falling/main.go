package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 256
	ScreenHeight = 240
)

type Game struct {
	sceneManager *SceneManager
	input        Input
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{})
	}

	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	pair1 = "apple"
	pair2 = "banana"
	pair3 = "cherry"
	pair4 = "orange"

	pairs = []string{pair1, pair2, pair3, pair4}

	fallingPair string
)

func init() {
	rand.Seed(time.Now().UnixNano())
	fallingPair = pairs[rand.Intn(len(pairs))]
}

func update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		fmt.Println("A key pressed")
	}

	return nil
}

func draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fallingPair)
}

func main() {
	if err := ebiten.RunGame(update, screenWidth, screenHeight, 1, "Word Match Game"); err != nil {
		panic(err)
	}
}
