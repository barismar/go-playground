package channel

import (
	"fmt"
	"time"
)

func numsToUnBufChan(nums []int) chan int {
	ch := make(chan int)
	go func() {
		for _, num := range nums {
			time.Sleep(2 * time.Second)
			ch <- num * 2
			fmt.Println("[blocking...]")
		}
		close(ch)
	}()
	return ch
}

func numsToBufChan(nums []int) chan int {
	ch := make(chan int, len(nums))
	go func() {
		for _, num := range nums {
			time.Sleep(2 * time.Second)
			ch <- num * 2
			fmt.Println("[unblocking...]")
		}
		close(ch)
	}()
	return ch
}

func Run() {
	fmt.Println("start...")

	// Unbuffered channel
	start := time.Now()
	nums := []int{1, 2, 3, 4, 5, 6}
	unBufCh := numsToUnBufChan(nums)

	for ubc := range unBufCh {
		fmt.Println("process unBufCh")
		fmt.Println("unBufCh:", ubc)
	}

	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("time for unbuffered channel:", diff)

	// Buffered channel
	start = time.Now()

	nums = []int{1, 2, 3, 4, 5, 6}
	bufCh := numsToBufChan(nums)

	for bc := range bufCh {
		fmt.Println("process bufCh")
		fmt.Println("bufCh:", bc)
	}

	end = time.Now()
	diff = end.Sub(start)
	fmt.Println("time for buffered channel:", diff)

	fmt.Println("exit...")
}
