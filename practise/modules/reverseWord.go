package modules

import "log"

func reverseWord(str string) string {
	if len(str) <= 1 {
		return str
	}

	lb, ub := 0, len(str)-1
	rns := []rune(str)
	for lb < ub {
		rns[lb], rns[ub] = rns[ub], rns[lb]
		lb++
		ub--
	}
	return string(rns)
}

func ReverseWord() {
	log.Printf("shivbaba: %s", reverseWord("shivbaba"))
	log.Printf("sandeep: %s", reverseWord("sandeep"))
	log.Printf("Single char s: %s", reverseWord("s"))
	log.Printf("Empty string: %s", reverseWord(""))
}
