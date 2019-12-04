package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type numbers []int

func newNumbers(from, to int) numbers {
	var ns numbers
	for i := from; i < to; i++ {
		ns = append(ns, i)
	}
	return ns
}

var d = []int{1000000, 100000, 10000, 1000, 100, 10}

func main() {
	from, to := mustRead("../input.txt")

	//ns := newNumbers(from, to)

	filtered := filterDoubles(from, to)
	filtered = filterDecreasingNumbers(filtered)

	fmt.Println("count:", len(filtered))
}

func filterDoubles(from, to int) numbers {
	var filtered numbers
	for i := from; i < to; i++ {
		ns := digits(i)
		if ns.hasDouble() {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

func filterDecreasingNumbers(ns numbers) numbers {
	var filtered numbers
	for _, i := range ns {
		ds := digits(i)
		if ds.isIncreasing() {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

func digits(n int) numbers {
	var ns numbers
	for _, i := range d {
		ns = append(ns, (n%i)/(i/10))
	}
	return ns
}

func (ns numbers) hasDouble() bool {
	var double bool
	for i, n := range ns {
		if next := i + 1; next < len(ns) {
			if ns[next] == n {
				double = true
				break
			}
		}
	}
	return double
}

func (ns numbers) isIncreasing() bool {
	increasing := true
	for i, n := range ns {
		if next := i + 1; next < len(ns) {
			if ns[next] < n {
				increasing = false
				break
			}
		}
	}
	return increasing
}

func mustRead(fn string) (from int, to int) {
	b, err := ioutil.ReadFile(fn)
	fatalOnError(err)

	splitted := bytes.Split(b, []byte("-"))

	from, err = strconv.Atoi(string(splitted[0]))
	fatalOnError(err)

	to, err = strconv.Atoi(string(splitted[1][:len(splitted[1])-1]))
	fatalOnError(err)

	return from, to
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
