package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sx := nextInt()
	sy := nextInt()
	tx := nextInt()
	ty := nextInt()

	if (sx+sy)%2 == 1 {
		sx--
	}
	if (tx+ty)%2 == 1 {
		tx--
	}

	dx := abs(sx - tx)
	dy := abs(sy - ty)
	ans := dy + max(0, (dx-dy)/2)
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
