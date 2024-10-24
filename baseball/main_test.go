package main

import "testing"
import "strings"
import "strconv"
import "fmt"

func TestBaseballNumber(t *testing.T) {
	numberOfTests := 100
	for i := 0; i < numberOfTests; i++ {
		number := makeBaseballNumber()

		split := strings.Split(number, "")
		intSplit := make([]int, len(split))
		for i, _ := range intSplit {
			if converted, err := strconv.Atoi(split[i]); err == nil {
				intSplit[i] = converted
			} else {
				fmt.Printf("%T, %v", converted, converted)
			}
		}

		if intSplit[0] == intSplit[1] || intSplit[0] == intSplit[2] || intSplit[1] == intSplit[2] {
			t.Errorf("error: %v", intSplit)
		}
	}
}
