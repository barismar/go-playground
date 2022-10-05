package main

import (
	"bufio"
	"fmt"
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
		var n int64
		scanf("%d\n", &n)
		solve(n)
	}
}

func solve(n int64) {
	total := n - 3
	ans := (total / 3) - 1

	printf("%d\n", ans)
}
