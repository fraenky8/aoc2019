package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

func fatalOnError(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

type Direction struct {
	Direction string
	Steps     int
}

type Position struct {
	X int
	Y int
}

func (p Position) distance() int {
	return Abs(p.X) + Abs(p.Y)
}

func Abs(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}

type Wire struct {
	Directions []Direction
	Positions  map[Position]struct{}
}

func (w *Wire) calculatePositions() {
	var x, y int

	for _, d := range w.Directions {
		if d.Direction == up {
			// +y positions
			for s := 0; s < d.Steps; s++ {
				y++
				w.Positions[Position{X: x, Y: y}] = struct{}{}
			}
		} else if d.Direction == right {
			// +x positions
			for s := 0; s < d.Steps; s++ {
				x++
				w.Positions[Position{X: x, Y: y}] = struct{}{}
			}
		} else if d.Direction == down {
			// -y positions
			for s := 0; s < d.Steps; s++ {
				y--
				w.Positions[Position{X: x, Y: y}] = struct{}{}
			}
		} else if d.Direction == left {
			// -x positions
			for s := 0; s < d.Steps; s++ {
				x--
				w.Positions[Position{X: x, Y: y}] = struct{}{}
			}
		} else {
			log.Fatalf("unknown direction %v", d)
		}
	}
}

func main() {
	wire1, wire2 := mustRead("../input.txt")

	// educated guess after reading the input :)
	if len(wire1.Positions) != len(wire2.Positions) {
		log.Fatal("both wires are not of equal length")
	}

	pos := intersect(wire1, wire2)

	fmt.Printf("Position %+v: distance: %d\n", pos, pos.distance())
}

func intersect(w1, w2 Wire) Position {

	w1.calculatePositions()
	w2.calculatePositions()

	pos := Position{X: math.MaxInt16, Y: math.MaxInt16}

	for p1 := range w1.Positions {
		if _, ok := w2.Positions[p1]; ok {
			if p1.distance() < pos.distance() {
				pos = p1
				fmt.Println(pos)
			}
		}
	}

	return pos
}

func mustRead(fn string) (Wire, Wire) {
	file, err := os.Open(fn)
	fatalOnError(err)
	defer file.Close()

	var (
		line  int
		wire1 Wire
		wire2 Wire
	)

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		line++
		directions := strings.Split(fs.Text(), ",")
		for _, dir := range directions {
			steps, err := strconv.Atoi(dir[1:])
			fatalOnError(err)
			direction := dir[:1]
			if line == 1 {
				wire1.Directions = append(wire1.Directions, Direction{
					Direction: direction,
					Steps:     steps,
				})
			} else {
				wire2.Directions = append(wire2.Directions, Direction{
					Direction: direction,
					Steps:     steps,
				})
			}
		}
	}

	wire1.Positions = make(map[Position]struct{})
	wire2.Positions = make(map[Position]struct{})

	return wire1, wire2
}
