package main

import "fmt"

func main() {
	fmt.Println(isAnagram("rot", "tor"))
}

func isAnagram(s, t string) bool {

	letters := make(map[rune]int)

	if len((s)) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	for _, char := range s {
		letters[char]++
	}

	for _, char := range t {
		letters[char]--
	}

	for _, value := range letters {
		if value != 0 {
			return false
		}
	}
	return true
}