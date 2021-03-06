package day1

import (
	"aoc2021/util"
	"fmt"
)

func Run() error {
	input, err := util.GetInputAsInts(1, false)
	if err != nil {
		return err
	}

	var prev int
	totDec := 0
	for i, line := range input {
		if i != 0 {
			if line > prev{
				totDec++
			}
		}
		
		prev = line
	}
	
	fmt.Printf("pt1: %d\n", totDec)

	totDec = 0
	for i := 0; i < len(input) - 2; i++ {
		cur := input[i] + input[i + 1] + input[i + 2]
		if i != 0 {
			if cur > prev {
				totDec++
			}
		}

		prev = cur
	}

	fmt.Printf("pt2: %d\n", totDec)

	return nil
}
