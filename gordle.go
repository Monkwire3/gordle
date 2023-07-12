package main

import (
	// "fmt"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	// "github.com/hajimehoshi/ebiten/v2/inpututil"

	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"math/rand"
)

const (
	title  string = "Gordle"
	width  int    = 435
	height int    = 600
	rows   int    = 6
	cols   int    = 5
)

var (
	fontSize        float64 = 24
	mplusNormalFont font.Face

	bkg       = color.White
	lightgrey = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	grey      = color.RGBA{0x77, 0x7c, 0x7e, 0xff}
	yellow    = color.RGBA{0xcd, 0xb3, 0x5d, 0xff}
	green     = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	fontColor = color.Black

	edge = false

	alphabet = "abcdefghijklmnopqrstuvwxyz"
	grid     [cols * rows]string
	dict     []string
	check    [cols * rows]int
	loc      int = 0
	won          = false
	answer   string
)

type Game struct {
	// For keyboard input
	runes []rune
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)

	if d == 1 {
		return true
	}

	if d >= delay && (d-delay)&interval == 0 {
		return true
	}

	return false

}

func (g *Game) Update() error {

}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	g := &Game{}

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Create the window
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	content, err := ioutil.ReadFile(("dict.txt"))
	if err != nil {
		log.Fatal(err)
	} else {
		dict = strings.Split(string(content), "\n")
	}

	rand.Seed(time.Now().UnixNano())
	answer = dict[rand.Intn(len(dict))]
	fmt.Printf(answer)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
