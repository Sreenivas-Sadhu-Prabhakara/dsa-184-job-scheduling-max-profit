package main

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		jobs []Job
		want int
	}{
		{[]Job{{1, 3, 50}, {2, 4, 10}, {3, 5, 40}, {3, 6, 70}}, 120}, // 50 + 70
		{[]Job{{1, 2, 4}, {3, 4, 3}, {2, 3, 2}}, 9},                  // all three fit
		{[]Job{{1, 5, 8}, {2, 3, 7}, {4, 6, 6}}, 13},                 // 7 + 6 beats 8
		{[]Job{{1, 10, 5}}, 5},
		{nil, 0},
		{[]Job{{1, 4, 10}, {1, 4, 20}}, 20}, // identical spans: take the better
	}
	for _, c := range cases {
		if got := MaxProfit(c.jobs); got != c.want {
			t.Errorf("MaxProfit(%v) = %d, want %d", c.jobs, got, c.want)
		}
	}
}
