package main

import "sort"

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	compte := 0
	fin := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < fin {
			compte++
		} else {
			fin = intervals[i][1]
		}
	}
	return compte
}

// func main() {
// 	println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
// 	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
// 	// il suffit de d'enlever qu'un seul intervalle, [1,3]
// 	println(Ft_non_overlap([][]int{{1, 2}, {2, 3}})) // resultat : 0
// 	println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}})) // resultat : 2
// }