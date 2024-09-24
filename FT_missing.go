package main

func Ft_missing(nums []int) int {
	n := len(nums)
	calcul := (n*(n+1)/2)
	actuel := 0

	for _, value := range nums {
		actuel += value
	}

	return calcul - actuel
}


// func main() {
// 	println(Ft_missing([]int{3, 1, 2})) // resultat : 0
// 	println(Ft_missing([]int{0, 1})) // resultat : 2
// 	println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8
// }