package main

import (
	"math"
)

func Ft_min_window(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Compte des caractères de t
	count := make(map[rune]int)
	for _, as := range t {
		count[as]++
	}

	windowCount := make(map[rune]int)
	left, right, formed := 0, 0, 0
	min, minLeft := math.MaxInt32, 0

	for right < len(s) {
		as := rune(s[right])
		windowCount[as]++

		// Vérification des caractères requis
		if windowCount[as] == count[as] {
			formed++
		}

		// Réduction de la fenêtre si tous les caractères sont présents
		for formed == len(count) {
			if right-left+1 < min {
				min, minLeft = right-left+1, left
			}
			windowCount[rune(s[left])]--
			if windowCount[rune(s[left])] < count[rune(s[left])] {
				formed--
			}
			left++
		}
		right++
	}

	if min == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft + min]
}

func main() {
	println(Ft_min_window("ADOBECODEBANC", "ABC")) // resultat : "BANC"
	println(Ft_min_window("a", "aa")) // resultat : ""
}