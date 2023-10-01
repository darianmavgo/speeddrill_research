package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/ebitenui/ebitenui"
	ebitenimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func init() {
	whiteImage.Fill(color.White)
}

const (
	screenWidth  = 640
	screenHeight = 480
	textSize     = 24
	gravity      = 1.0
	fontFile     = "/Users/darianhickman/Documents/Consolas.ttf" // Update with your TTF font file path
)

var (
	fontFace   font.Face
	text       string
	yPos       float64
	whiteImage = ebiten.NewImage(3, 3)
)

type Game struct {
	ui *ebitenui.UI
	//This parameter is so you can keep track of the textInput widget to update and retrieve
	//its values in other parts of your game
	standardTextInput *widget.TextInput
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return fmt.Errorf("game closed by user")
	}

	yPos += gravity

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, yPos)

	// ebitenutil.DebugPrint(screen, "F=ma")
	// ebitenutil.DebugPrintAt(screen, text, 0, int(yPos)+textSize)
	// update the UI
	// Additional keys to manage focus
	if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) {
		g.ui.ChangeFocus(ebitenui.FOCUS_PREVIOUS)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) {
		g.ui.ChangeFocus(ebitenui.FOCUS_NEXT)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnd) {
		if g.ui.GetFocusedWidget() == g.standardTextInput {
			fmt.Println("standardTextInput selected")
		}
	}
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
	d := &font.Drawer{
		Dst:  screen,
		Src:  image.White,
		Face: fontFace,
		Dot:  fixed.P(0, int(yPos)+2*textSize),
	}
	d.DrawString(text)
	// ebitenutil.DebugPrint(screen, msg)
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadFont() font.Face {
	fontData, err := os.ReadFile(fontFile)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := truetype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	return truetype.NewFace(tt, &truetype.Options{
		Size:    textSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func main() {
	fontFace = loadFont()

	text = "F=ma"
	yPos = 0
	game.standardTextInput = widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			//Set the layout information to center the textbox in the parent
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),

		//Set the Idle and Disabled background image for the text input
		//If the NineSlice image has a minimum size, the widget will use that or
		// widget.WidgetOpts.MinSize; whichever is greater
		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:     ebitenimage.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
			Disabled: ebitenimage.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
		}),

		//Set the font face and size for the widget
		widget.TextInputOpts.Face(face),

		//Set the colors for the text and caret
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          color.NRGBA{254, 255, 255, 255},
			Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Caret:         color.NRGBA{254, 255, 255, 255},
			DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		}),

		//Set how much padding there is between the edge of the input and the text
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),

		//Set the font and width of the caret
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(face, 2),
		),

		//This text is displayed if the input is empty
		widget.TextInputOpts.Placeholder("Standard Textbox"),

		//This is called when the user hits the "Enter" key.
		//There are other options that can configure this behavior
		widget.TextInputOpts.SubmitHandler(func(args *widget.TextInputChangedEventArgs) {
			fmt.Println("Text Submitted: ", args.InputText)
		}),

		//This is called whenver there is a change to the text
		widget.TextInputOpts.ChangedHandler(func(args *widget.TextInputChangedEventArgs) {
			fmt.Println("Text Changed: ", args.InputText)
		}),
	)

	ebiten.SetWindowSize(screenWidth, screenHeight)

	ebiten.SetWindowTitle("Falling Text Input - F=ma")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
