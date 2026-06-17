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
		dir:      Point{x: -1, y: 0},
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
		termbox.SetCell(x, g.height+1, '─', termbox.ColorCyan, termbox.ColorDefault)
	}
	for y := 1; y <= g.height; y++ {
		termbox.SetCell(0, y, '│', termbox.ColorCyan, termbox.ColorDefault)
		termbox.SetCell(g.width+1, y, '│', termbox.ColorCyan, termbox.ColorDefault)
	}

	termbox.SetCell(0, 0, '┌', termbox.ColorCyan, termbox.ColorDefault)
	termbox.SetCell(g.width+1, 0, '┐', termbox.ColorCyan, termbox.ColorDefault)
	termbox.SetCell(0, g.height+1, '└', termbox.ColorCyan, termbox.ColorDefault)
	termbox.SetCell(g.width+1, g.height+1, '┘', termbox.ColorCyan, termbox.ColorDefault)

	text := fmt.Sprintf("Score: %d Level: %d", g.score, g.level)
	for i, ch := range text {
		termbox.SetCell(i+2, 1, ch, termbox.ColorYellow, termbox.ColorDefault)
	}

	head := g.snake[0]
	screenX := head.x
	screenY := head.y + 1

	headToRune := g.dir.ToRune()

	termbox.SetCell(screenX, screenY, headToRune, termbox.ColorGreen, termbox.ColorDefault)
	if len(g.snake) > 1 {
		for _, body := range g.snake[1:] {
			termbox.SetCell(body.x, body.y+1, '○', termbox.ColorGreen, termbox.ColorDefault)
		}
	}

	err = termbox.Flush()
	if err != nil {
		return
	}
}

func (p Point) ToRune() rune {
	x := p.x
	y := p.y

	switch {
	case x > 0:
		return '▶'
	case x < 0:
		return '◀'
	case y > 0:
		return '▼'
	case y < 0:
		return '▲'
	default:
		return '●'
	}
}

func (g *Game) handleInput(ev termbox.Event) {
	eventType := ev.Type
	if eventType != termbox.EventKey {
		return
	}

	switch {
	case ev.Key == termbox.KeyEsc || ev.Ch == 'q' || ev.Ch == 'Q' || ev.Ch == 'й' || ev.Ch == 'Й':
		select {
		case <-g.quit:
		default:
			close(g.quit)
		}
		return

	case (ev.Key == termbox.KeyArrowUp) || ev.Ch == 'w' || ev.Ch == 'W' || ev.Ch == 'ц' || ev.Ch == 'Ц':
		if g.dir.y != 1 {
			g.dir = Point{0, -1}
		}

	case ev.Key == termbox.KeyArrowRight || ev.Ch == 'd' || ev.Ch == 'D' || ev.Ch == 'в' || ev.Ch == 'В':
		if g.dir.x != -1 {
			g.dir = Point{1, 0}
		}

	case ev.Key == termbox.KeyArrowDown || ev.Ch == 's' || ev.Ch == 'S' || ev.Ch == 'ы' || ev.Ch == 'Ы':
		if g.dir.y != -1 {
			g.dir = Point{0, 1}
		}

	case ev.Key == termbox.KeyArrowLeft || ev.Ch == 'a' || ev.Ch == 'A' || ev.Ch == 'ф' || ev.Ch == 'Ф':
		if g.dir.x != 1 {
			g.dir = Point{-1, 0}
		}
	}
}

func (g *Game) isOnSnake(p Point) bool {
	for _, s := range g.snake {
		if s.x == p.x && s.y == p.y {
			return true
		}
	}

	return false
}

func (g *Game) isOnMalware(p Point) bool {
	for _, m := range g.malware {
		if m.x == p.x && m.y == p.y {
			return true
		}
	}

	return false
}

func (g *Game) isOutOfBounds(p Point) bool {
	return p.x < 1 || p.x > g.width || p.y < 1 || p.y > g.height
}

func main() {
	err := termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()

	eventCh := make(chan termbox.Event)

	game := NewGame(40, 20)
	game.draw()

	go func() {
		for {
			eventCh <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventCh:
			game.handleInput(ev)
			game.draw()
		case <-game.quit:
			return
		}
	}
}
