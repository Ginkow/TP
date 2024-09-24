package main

func Ft_max_substring(s string) int {
	a := make(map[rune]int)
	start, max := 0, 0

	for fin, as := range s {
		if last, found := a[as]; found && last >= start {
			start = last + 1
		}
		a[as] = fin
		if current := fin - start + 1; current > max {
			max = current
		}
	}

	return max
}

// func main() {
// 	println(Ft_max_substring("abcabcbb")) // resultat : 3
// 	// "abc" est la plus grande sous chaine composé de caractère diffèrent
// 	println(Ft_max_substring("bbbbb")) // resultat : 1
// 	// "b" est la plus grande sous chaine
// }
