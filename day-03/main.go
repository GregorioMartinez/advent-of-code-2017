package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(distAccessPort(265149))
}

func distAccessPort(i float64) int {

	if i == 1 {
		return 0
	}

	n := int(i)

	l := int(math.Ceil(math.Sqrt(i)))
	if l%2 == 0 {
		l = l + 1
	}

	mid := l / 2
	steps := mid * 2
	corner := l * l
	swap := true

	for k := corner; k > n; k-- {
		if steps == mid || steps == mid*2 {
			swap = !swap
		}

		if !swap {
			steps--
		} else {
			steps++
		}

	}
	return steps
}
