package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(format string, args ...interface{})  { fmt.Fscanf(reader, format, args...) }
func printf(format string, args ...interface{}) { fmt.Fprintf(writer, format, args...) }

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var n int
		var a int
		minVal := math.MaxInt32

		scanf("%d\n", &n)
		var data []int

		for j := 0; j < n; j++ {
			if j == n-1 {
				scanf("%d\n", &a)
			} else {
				scanf("%d ", &a)
			}

			if minVal > a {
				minVal = a
			}

			data = append(data, a)
		}

		solve(data, minVal)
	}
}

func solve(data []int, minVal int) {
	var answer int

	printf("%d\n", answer)
}
