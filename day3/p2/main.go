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

type Wire struct {
	Directions []Direction
	Positions  map[Position]int
}

type Intersections map[Position]int

func (w *Wire) calculatePositions() {
	var (
		x, y  int
		steps = 1
	)

	for _, d := range w.Directions {
		if d.Direction == up {
			// +y positions
			for s := 0; s < d.Steps; s++ {
				y++
				w.Positions[Position{X: x, Y: y}] = steps
				steps++
			}
		} else if d.Direction == right {
			// +x positions
			for s := 0; s < d.Steps; s++ {
				x++
				w.Positions[Position{X: x, Y: y}] = steps
				steps++
			}
		} else if d.Direction == down {
			// -y positions
			for s := 0; s < d.Steps; s++ {
				y--
				w.Positions[Position{X: x, Y: y}] = steps
				steps++
			}
		} else if d.Direction == left {
			// -x positions
			for s := 0; s < d.Steps; s++ {
				x--
				w.Positions[Position{X: x, Y: y}] = steps
				steps++
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

	pos, steps := solve(wire1, wire2)

	fmt.Printf("Position %+v: steps: %d\n", pos, steps)
}

func solve(w1, w2 Wire) (Position, int) {
	intersections := intersect(w1, w2)
	return findMinimum(intersections)
}

func intersect(w1, w2 Wire) Intersections {

	w1.calculatePositions()
	w2.calculatePositions()

	intersections := make(Intersections)

	for p1, s1 := range w1.Positions {
		if s2, ok := w2.Positions[p1]; ok {
			intersections[p1] = s1 + s2
		}
	}

	return intersections
}

func findMinimum(intersections Intersections) (Position, int) {
	var (
		steps    = math.MaxInt16
		position Position
	)

	for pos, s := range intersections {
		if s < steps {
			steps = s
			position = pos
		}
	}

	return position, steps
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

	wire1.Positions = make(map[Position]int)
	wire2.Positions = make(map[Position]int)

	return wire1, wire2
}
