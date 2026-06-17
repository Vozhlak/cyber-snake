package main

import "fmt"

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

func main() {
	game := NewGame(20, 40)

	fmt.Printf(
		"Игра создана: поле %dx%d, змейка в (%d, %d), направление вправо, уровень %d\n",
		game.width,
		game.height,
		game.snake[0].x,
		game.snake[0].y,
		game.level,
	)
}
