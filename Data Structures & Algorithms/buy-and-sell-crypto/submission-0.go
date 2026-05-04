func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = min(dp[i-1], prices[i])
	}
	var maxProfit int
	maxProfit = 0
	for i := 0; i < len(dp); i++ {
		profit := prices[i] - dp[i]
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	
	return maxProfit
}
