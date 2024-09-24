package main

func Ft_coin(coins []int, amount int) int {
	piece := make([]int, amount + 1)

	for i := range piece {
		piece[i] = amount + 1
	}
	piece[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				piece[i] = min(piece[i], piece[i-coin]+1)
			}
		}
	}

	if piece[amount] > amount {
		return -1
	}
	return piece[amount]
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
	

// func main() {
// 	println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
// 	println(Ft_coin([]int{2}, 3)) // resultat : -1
// 	println(Ft_coin([]int{1}, 0)) // resultat : 0
// }