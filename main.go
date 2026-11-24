package main

import (
	"fmt"
	"sort"
)

// Job runs over [Start, End) and pays Profit.
type Job struct {
	Start, End, Profit int
}

// MaxProfit returns the best total profit from non-overlapping jobs
// (weighted interval scheduling): sort by end time; each job either joins
// the best schedule ending at its latest compatible predecessor, or sits out.
func MaxProfit(jobs []Job) int {
	n := len(jobs)
	if n == 0 {
		return 0
	}
	sorted := append([]Job{}, jobs...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].End < sorted[j].End })
	ends := make([]int, n)
	for i, j := range sorted {
		ends[i] = j.End
	}
	dp := make([]int, n+1) // dp[i]: best profit using the first i jobs
	for i := 1; i <= n; i++ {
		job := sorted[i-1]
		// Latest job whose end ≤ this job's start (touching is compatible).
		p := sort.SearchInts(ends[:i-1], job.Start+1)
		take := dp[p] + job.Profit
		skip := dp[i-1]
		if take > skip {
			dp[i] = take
		} else {
			dp[i] = skip
		}
	}
	return dp[n]
}

func main() {
	jobs := []Job{{1, 3, 50}, {2, 4, 10}, {3, 5, 40}, {3, 6, 70}}
	fmt.Println(MaxProfit(jobs))
}
