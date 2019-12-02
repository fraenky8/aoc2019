package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func fatalOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

const expected = 19690720

func main() {
	codes := mustRead("../input.txt")

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			output := run(n, v, codes)
			if output[0] == expected {
				fmt.Printf("noun: %v, verb: %v\n100 * noun + verb = %v\n", n, v, 100*n+v)
				return
			}
		}
	}

	fmt.Println("could not find output!")
}

func mustRead(fn string) []int {
	input, err := ioutil.ReadFile(fn)
	fatalOnError(err, "could not open input")
	byteCodes := bytes.Split(input, []byte(","))

	var codes []int
	for _, b := range byteCodes {
		n, _ := strconv.Atoi(string(b))
		codes = append(codes, n)
	}

	if len(byteCodes) != len(codes) {
		log.Fatal("reading input failed")
	}

	return codes
}

const (
	addition       = 1
	multiplication = 2
)

func run(noun, verb int, c []int) []int {
	var codes = make([]int, len(c))
	copy(codes, c)

	codes[1] = noun
	codes[2] = verb

	for i := 0; i < len(codes); i += 4 {
		op := codes[i]
		if op == 99 {
			break
		}
		in1 := codes[codes[i+1]]
		in2 := codes[codes[i+2]]

		if op == addition {
			codes[codes[i+3]] = in1 + in2
		} else if op == multiplication {
			codes[codes[i+3]] = in1 * in2
		} else {
			return codes
		}
	}
	return codes
}
