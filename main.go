package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
	"log"
)

type Goober struct {
	x         float64
	y         float64
	sprite    *ebiten.Image
	direction bool
}

var goober Goober

const (
	windowWidth  = 640
	windowHeight = 480
)

var (
	gophersImage *ebiten.Image
	imageX       = 0.0
	imageY       = 0.0
	rightSpeed   = 10.0
	leftSpeed    = 10.0
	upSpeed      = 10.0
	downSpeed    = 10.0
)

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error { // Updates every frame/tick.
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) { // Updates every frame/tick.
	// Determines the movement of the player.
	playerMove(g.keys)

	// Inserts enemy.
	gooberAI(&goober)
	enemyOp := &ebiten.DrawImageOptions{}
	enemyOp.GeoM.Translate(goober.x, goober.y)
	screen.DrawImage(goober.sprite, enemyOp)

	// Moves the player.
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(imageX, imageY)
	screen.DrawImage(gophersImage, playerOp)

	// Displays the coordinates of the player.
	displayCoordinate(screen, imageX, imageY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func main() {
	img, _, err := ebitenutil.NewImageFromFile("goober_resized.png")
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)

	enemy, _, err := ebitenutil.NewImageFromFile("goober_resized.png")
	if err != nil {
		log.Fatal(err)
	}

	goober = Goober{
		50,
		50,
		ebiten.NewImageFromImage(enemy),
		true,
	}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Goober simulator")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func playerMove(pressedKeys []ebiten.Key) {
	for _, key := range pressedKeys {
		if key == ebiten.KeyArrowRight {
			if worldBorderStatus(imageX+rightSpeed, imageY) {
				imageX += rightSpeed
			} else {
				imageX = windowWidth
			}
		} else if key == ebiten.KeyArrowLeft {
			if worldBorderStatus(imageX-leftSpeed, imageY) {
				imageX -= leftSpeed
			} else {
				imageX = 0
			}
		} else if key == ebiten.KeyArrowUp {
			if worldBorderStatus(imageX, imageY-upSpeed) {
				imageY -= upSpeed
			} else {
				imageY = 0
			}
		} else if key == ebiten.KeyArrowDown {
			if worldBorderStatus(imageX, imageY+downSpeed) {
				imageY += downSpeed
			} else {
				imageY = windowHeight
			}
		}
	}
}

func worldBorderStatus(x, y float64) bool {
	// Checks if a move would pu the character out of bound.
	if x < 0 || x > windowWidth || y < 0 || y > windowHeight {
		return false
	}
	return true
}

func displayCoordinate(screen *ebiten.Image, x, y float64) {
	coordinate := fmt.Sprintf("(%.2f, %.2f)", x, y)
	ebitenutil.DebugPrint(screen, coordinate)
}

func gooberAI(goober *Goober) {
	if goober.x > windowWidth {
		goober.direction = false
	} else if goober.x < 0 {
		goober.direction = true
	}
	fmt.Println(goober.direction)
	fmt.Println(goober.x)
	if goober.direction {
		goober.x += 5
	} else {
		goober.x -= 5
	}
}
