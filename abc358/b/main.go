package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := nextInt()
	a := nextInt()
	t := nextInts(n)
	ans := make([]int, n+1)
	for i := 0; i < n; i++ {
		ans[i+1] = max(ans[i], t[i]) + a
	}
	strAns := make([]string, n+1)
	for i := 0; i < n; i++ {
		strAns[i] = strconv.Itoa(ans[i+1])
	}
	fmt.Println(strings.Join(strAns, "\n"))
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
		nums[i] = nextInt()
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
