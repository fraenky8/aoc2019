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
	filtered = filterLargeGroups(filtered)

	fmt.Println("count:", len(filtered))
	fmt.Println(filtered)
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

func filterLargeGroups(ns numbers) numbers {
	var filtered numbers
	for _, i := range ns {
		ds := digits(i)
		if !ds.isInLargerGroup() {
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

func (ns numbers) isInLargerGroup() bool {
	groups := make([]int, 10, 10)
	for _, n := range ns {
		groups[n]++
	}
	for i, v := range groups {
		if v == 1 {
			groups[i] = 0
		}
	}

	var idx []int
	for i := range groups {
		if groups[i] > 0 {
			idx = append(idx, i)
		}
	}

	if len(idx) == 1 {
		return groups[idx[0]] != 2
	}

	if len(idx) == 2 {
		if groups[idx[0]] == 3 && groups[idx[1]] == 3 {
			return true
		}
	}

	// 3 groups always false
	return false
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
