package main

func Ft_profit(prix []int) int {
	if len(prix) < 2 {
		return 0
	}
	min := prix[0]
	max_profit := 0
	for i := 1; i < len(prix); i++ {
		if prix[i] < min {
			min = prix[i]
		} else {
			max_profit = max(max_profit, prix[i]-min)
		}
	}
	return max_profit
}

// func main() {
// 	println(Ft_profit([]int{7,1,5,3,6,4})) // resultat : 5
// 	// si on achète au jour 1, nous payons 1,
// 	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
// 	println(Ft_profit([]int{7,6,4,3,1})) // resultat : 0
// }