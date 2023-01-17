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
		bucket := make(map[int]int)
		var itemsCount, price, sum int

		fmt.Fscan(in, &itemsCount)
		for i := 0; i < itemsCount; i++ {
			fmt.Fscan(in, &price)
			bucket[price]++
		}

		for price, cnt := range bucket {
			effectiveCount := cnt - (cnt / 3)
			sum += price * effectiveCount
		}
		fmt.Fprintln(out, sum)
	}
}

/*
6
12
2 2 2 2 2 2 2 3 3 3 3 3
12
2 3 2 3 2 2 3 2 3 2 2 3
1
10000
9
1 2 3 1 2 3 1 2 3
6
10000 10000 10000 10000 10000 10000
6
300 100 200 300 200 300
*/
