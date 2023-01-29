package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount int
	fmt.Fscan(in, &testsCount)

	for t := 0; t < testsCount; t++ {

		var sn, sm, n, m int
		fmt.Fscan(in, &sn, &sm)

		sheet := make([]*[]int, sn)
		for n = 0; n < sn; n++ {
			nc := make([]int, sm)
			for m = 0; m < sm; m++ {
				fmt.Fscan(in, &nc[m])
			}
			sheet[n] = &nc
		}

		var clicksCount, clickIdx int
		fmt.Fscan(in, &clicksCount)
		var clicks []int
		for c := 0; c < clicksCount; c++ {
			fmt.Fscan(in, &clickIdx)
			clicks = append(clicks, clickIdx-1)
		}

		// clicks
		for _, clickIdx := range clicks {
			colMap := make(map[int][]*[]int, sn)
			for n = 0; n < sn; n++ {
				nc := *sheet[n]
				cellValue := nc[clickIdx]
				colMap[cellValue] = append(colMap[cellValue], &nc)
			}
			var values []int
			for value := range colMap {
				values = append(values, value)
			}
			sort.Ints(values)
			var rowsSorted []*[]int
			for _, value := range values {
				rowsSorted = append(rowsSorted, colMap[value]...)
			}

			var sheetNext []*[]int
			for _, row := range rowsSorted {
				sheetNext = append(sheetNext, row)
			}
			sheet = sheetNext
		}

		// out
		for n = 0; n < sn; n++ {
			nc := *sheet[n]
			for m = 0; m < sm; m++ {
				fmt.Fprint(out, nc[m], " ")
			}
			fmt.Fprint(out, "\n")
		}
		fmt.Fprintln(out, "")
	}

}
