package best_time_to_buy_and_sell_stock_II

// BestTimeToBuyAndSellStockII [time complexity:O(n),auxiliary space:O(1)]
func BestTimeToBuyAndSellStockII(prices []int) int {
	n := len(prices)
	maxProfit := 0
	for i := range n - 1 {
		if prices[i+1] > prices[i] {
			maxProfit += prices[i+1] - prices[i]
		}
	}
	return maxProfit
}
