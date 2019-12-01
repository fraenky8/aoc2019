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

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fuel int

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		s := strings.TrimSpace(fs.Text())
		m, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fuel += int(math.Floor(float64(m/3)) - 2)
	}

	fmt.Println(fuel)
}
