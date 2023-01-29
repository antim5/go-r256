package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount int
	fmt.Fscan(in, &testsCount)

	for t := 0; t < testsCount; t++ {
		rosterBySkill := make(map[int][]int)
		var itemsCount, skillLevel, idx, skillDelta int
		var roster []int
		rosterOut := make(map[int]bool)

		fmt.Fscan(in, &itemsCount)
		for idx = 0; idx < itemsCount; idx++ {
			fmt.Fscan(in, &skillLevel)
			roster = append(roster, skillLevel)
			rosterBySkill[skillLevel] = append(rosterBySkill[skillLevel], idx)
		}

		if t != 0 {
			fmt.Fprintln(out, "")
		}
		for idx = 0; idx < itemsCount; idx++ {
			if rosterOut[idx] == true {
				continue
			}
			skillLevel = roster[idx]
			skillDelta = 0
			rosterBySkill[skillLevel] = rosterBySkill[skillLevel][1:]
		sLoop:
			for {
				if skillDelta > 100 {
					break sLoop
				}
				var leftIdx, rightIdx int
				if left := rosterBySkill[skillLevel-skillDelta]; len(left) != 0 {
					leftIdx = left[0]
				}
				if right := rosterBySkill[skillLevel+skillDelta]; len(right) != 0 {
					rightIdx = right[0]
				}
				switch {
				case leftIdx != idx && !rosterOut[leftIdx] && (rightIdx == 0 || (leftIdx != 0 && leftIdx <= rightIdx)):
					fmt.Fprintln(out, idx+1, leftIdx+1)
					rosterOut[idx] = true
					rosterOut[leftIdx] = true
					rosterBySkill[skillLevel-skillDelta] = rosterBySkill[skillLevel-skillDelta][1:]
					break sLoop
				case rightIdx != idx && !rosterOut[rightIdx]:
					fmt.Fprintln(out, idx+1, rightIdx+1)
					rosterOut[idx] = true
					rosterOut[rightIdx] = true
					rosterBySkill[skillLevel+skillDelta] = rosterBySkill[skillLevel+skillDelta][1:]
					break sLoop
				default:
					skillDelta++
					continue
				}
			}
		}

	}
}

/*
3
6
2 1 3 1 1 4
2
5 5
8
1 4 2 5 4 2 6 3

*/
