package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
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

func nextInts() []int {
	s := strings.Split(nextLine(), " ")
	nums := make([]int, len(s))
	for i, v := range s {
		nums[i], _ = strconv.Atoi(v)
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
