package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	n := nextInt()
	a := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		a[i] = nextInt()
	}
	ans := 0
	for i := 0; i < n*2-2; i++ {
		if a[i] == a[i+2] {
			ans++
		}
	}
	fmt.Println(ans)
}

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextLine())
	return i
}

func nextInts(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(nextLine())
	}
	return nums
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
