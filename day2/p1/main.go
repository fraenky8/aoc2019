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

func main() {
	codes := mustRead("../input.txt")

	// To do this, before running the program,
	// - replace position 1 with the value 12
	// - and replace position 2 with the value 2.
	codes[1] = 12
	codes[2] = 2

	output := restore(codes)

	fmt.Printf("%v\n", output)
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

func restore(codes []int) []int {
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
			log.Fatalf("invalid opcode %v detected\n", op)
		}
	}
	return codes
}
