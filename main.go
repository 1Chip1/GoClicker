package main

import (
	"fmt"
	"strconv"
	"log"
	_"image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_"github.com/hajimehoshi/ebiten/v2/text"
)

// the consts 
const (
	screenWidth = 300
	screenHeight = 300
)

// the variables 
var (
	Score = 0
	ScoreText = ""
	win = false
)


// is the mouse button just pressed?
func IsMouseButtonJustPressed(button ebiten.MouseButton) bool {
	return true
}

// creating the game 
type Game struct {}


// window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}


// update function
func (g *Game) Update() error {	
	// if the mouse just pressed add score to the Score variable
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		Score = Score + 1
		ScoreText = strconv.Itoa(Score)
		fmt.Println(ScoreText)
	}

	if Score >= 100 {
		win = true
	}
	return nil
}


// draw function for the text and the Ebitengin logo
func (g *Game) Draw(screen *ebiten.Image) {
	if win != true{
		// printing the Score to the screen
		ebitenutil.DebugPrint(screen, ScoreText)
	}

	// creating the logo 
	logo, _, err := ebitenutil.NewImageFromFile("logo.png")
	if err != nil {
		log.Fatal(err)
	}
	op := &ebiten.DrawImageOptions{}

	// drawing the logo
	screen.DrawImage(logo, op)

	if win == true {
		ebitenutil.DebugPrint(screen, "You Win!!")
	}
}


// main function to start the game!
func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GoClicker")


	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}