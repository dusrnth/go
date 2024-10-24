package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	bbNum := makeBaseballNumber()
	fmt.Println(bbNum)
}

func makeBaseballNumber() string {
	var s = make([]int, 3)
	for n := range 3 {
		var i int
		for i == 0 {
			i = rand.Intn(9)
			fmt.Println(i)
		}

		for _, number := range s {
			if number != 0 && number == i {
				for i == 0 {
					i = rand.Intn(9)
				}
			}
		}

		s[n] = i
	}

	return convertToString(s)
}

func convertToString(s []int) string {
	str := make([]string, len(s))
	for i, number := range s {
		str[i] = strconv.Itoa(number)
	}
	return strings.Join(str, "")
}
