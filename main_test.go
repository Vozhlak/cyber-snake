package main

import "testing"

func TestIsOnSnake(t *testing.T) {
	g := &Game{
		snake: []Point{
			{20, 10},
			{19, 10},
			{18, 10},
		},
	}

	tests := []struct {
		name string
		p    Point
		want bool
	}{
		{"head", Point{20, 10}, true},
		{"body", Point{19, 10}, true},
		{"miss", Point{5, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := g.isOnSnake(tt.p)
			if got != tt.want {
				t.Errorf("isOnSnake(%v) = %v, want %v", tt.p, got, tt.want)
			}
		})
	}
}

func TestIsOnMalware(t *testing.T) {
	g := &Game{
		malware: []Point{
			{7, 7},
			{8, 8},
		},
	}

	tests := []struct {
		name string
		p    Point
		want bool
	}{
		{"first", Point{7, 7}, true},
		{"second", Point{8, 8}, true},
		{"miss", Point{5, 5}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := g.isOnMalware(tt.p)
			if got != tt.want {
				t.Errorf("isOnMalware(%v) = %v, want %v", tt.p, got, tt.want)
			}
		})
	}
}

func TestIsOutOfBounds(t *testing.T) {
	g := &Game{
		width:  40,
		height: 20,
	}

	tests := []struct {
		name string
		p    Point
		want bool
	}{
		{"left outside", Point{0, 5}, true},
		{"top outside", Point{5, 0}, true},
		{"right outside", Point{41, 5}, true},
		{"bottom outside", Point{5, 21}, true},

		{"top-left inside", Point{1, 1}, false},
		{"top-right inside", Point{40, 1}, false},
		{"bottom-left inside", Point{1, 20}, false},
		{"bottom-right inside", Point{40, 20}, false},
		{"middle inside", Point{20, 10}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := g.isOutOfBounds(tt.p)
			if got != tt.want {
				t.Errorf("isOutOfBounds(%v) = %v, want %v", tt.p, got, tt.want)
			}
		})
	}
}
