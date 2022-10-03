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

	var hi string
	scanf("%s\n", &hi)
	printf("%s\n", hi)
}
