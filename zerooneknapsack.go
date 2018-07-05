package learngo

// max returns max of two input integer
func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

/*
KnapSack returns max total value of items can be put in a box
given the weight capacity, weight and value of each item

w = max weight allowed
wt = weight table for each item
vt = value table for each item
 WT | VT
----|----
 1  |  1
 3  |  4
 4  |  5
 5  |  7

w = 7
 WT | VT | 0 1 2 3 4 5 6 7 = w
---------|------------------------------
  1 | 1  | 0 1 1 1 1 1 1 1
  3 | 4  | 0 1 1 4 5 5 5 5
  4 | 5  | 0 1 1 4 5 6 6 9
  5 | 7  | 0 1 1 4 5 7 8 9
*/
func KnapSack(w int64, wt []int64, vt []int64) int64 {
	wtLen := len(wt)
	vtLen := len(vt)

	if wtLen != vtLen {
		return -1
	}

	if wtLen == 0 {
		return 0
	}

	// Inital memorization of [n+1][w+1] to 0
	dp := make([][]int64, wtLen+1)
	for y := range dp {
		dp[y] = make([]int64, w+1)
	}

	for i := int64(1); i <= int64(wtLen); i = i + 1 {
		for j := int64(1); j <= w; j = j + 1 {
			if wt[i-1] <= j {
				dp[i][j] = max(vt[i-1]+dp[i-1][j-wt[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[wtLen][w]
}
