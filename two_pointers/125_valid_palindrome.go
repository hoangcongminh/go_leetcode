package main

import "unicode"

func isPalindrome(s string) bool {
	lo := 0
	hi := len(s) - 1
	arr := []rune(s)

	for lo < hi {
		loChar := unicode.ToLower(arr[lo])
		hiChar := unicode.ToLower(arr[hi])

		if !isLetterOrDigit(loChar) {
			lo++
			continue
		}

		if !isLetterOrDigit(hiChar) {
			hi--
			continue
		}

		if loChar != hiChar {
			return false
		}

		lo++
		hi--
	}

	return true
}

func isLetterOrDigit(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}
