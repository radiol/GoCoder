package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	n := nextInt()
	m := nextInt()
	a := nextInts(n)
	b := nextInts(m)
	sort.Ints(a)
	sort.Ints(b)
	j := 0
	ans := 0
	for i := 0; i < m; i++ {
		for j < n && b[i] > a[j] {
			j++
		}
		if j == n {
			fmt.Println(-1)
			return
		}
		ans += a[j]
		j++
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
