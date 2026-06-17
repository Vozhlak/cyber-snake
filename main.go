package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type Point struct {
	x int
	y int
}

type Game struct {
	snake    []Point
	food     Point
	malware  []Point
	dir      Point
	score    int
	level    int
	gameOver bool
	width    int
	height   int
	quit     chan struct{}
}

func NewGame(width, height int) *Game {
	return &Game{
		snake:    []Point{{x: width / 2, y: height / 2}},
		food:     Point{x: 4, y: 4},
		malware:  make([]Point, 0),
		dir:      Point{x: 1, y: 0},
		score:    0,
		level:    1,
		gameOver: false,
		width:    width,
		height:   height,
		quit:     make(chan struct{}),
	}
}

func (g *Game) draw() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return
	}

	for x := 1; x <= g.width; x++ {
		termbox.SetCell(x, 0, '─', termbox.ColorCyan, termbox.ColorDefault)
		termbox.SetCell(x, g.height+2, '─', termbox.ColorCyan, termbox.ColorDefault)
	}
	for y := 1; y <= g.height+1; y++ {
		termbox.SetCell(0, y, '│', termbox.ColorCyan, termbox.ColorDefault)
		termbox.SetCell(g.width+1, y, '│', termbox.ColorCyan, termbox.ColorDefault)
	}

	termbox.SetCell(0, 0, '┌', termbox.ColorCyan, termbox.ColorDefault)
	termbox.SetCell(g.width+1, 0, '┐', termbox.ColorCyan, termbox.ColorDefault)
	termbox.SetCell(0, g.height+2, '└', termbox.ColorCyan, termbox.ColorDefault)
	termbox.SetCell(g.width+1, g.height+2, '┘', termbox.ColorCyan, termbox.ColorDefault)

	text := fmt.Sprintf("Score: %d Level: %d", g.score, g.level)
	for i, ch := range text {
		termbox.SetCell(i+2, 1, ch, termbox.ColorYellow, termbox.ColorDefault)
	}

	head := g.snake[0]
	screenX := head.x
	screenY := head.y + 1

	termbox.SetCell(screenX, screenY, '█', termbox.ColorGreen, termbox.ColorDefault)

	err = termbox.Flush()
	if err != nil {
		return
	}
}

func main() {

	err := termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()

	game := NewGame(40, 20)
	game.draw()

	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if event.Key == termbox.KeyEsc || event.Ch == 'q' {
				break
			}
		}
	}
}
