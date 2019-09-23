package main

func main() {

}

func CheckUnique(word string) bool {
	unique := true
	var seen = make(map[rune]bool)
	for _, char := range word {
		_, found := seen[char]
		if found {
			return false
		}
		seen[char] = true
	}
	return unique
}
