func isAnagram(first string, second string) bool {
	var set map[rune]int = make(map[rune]int)

	for _, f := range first {
		set[f] += 1
	}

	for _, s := range second {
		if _, ok := set[s]; !ok {
			return false
		}

		if set[s] == 1 {
			delete(set, s)
			continue
		}

		set[s] -= 1

	}


	return len(set) == 0
}
