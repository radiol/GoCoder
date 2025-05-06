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
	m := nextInt()
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		ss[i] = nextLine()
	}
	bits := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ss[i][j] == 'o' {
				bits[i] |= 1 << j
			}
		}
	}
	ans := 1 << 30
	for b := 0; b < (1 << n); b++ {
		p := 0
		cnt := 0
		for i := 0; i < n; i++ {
			if b>>i&1 == 1 {
				p |= bits[i]
				cnt++
			}
		}
		if p == (1<<m)-1 {
			ans = min(ans, cnt)
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
