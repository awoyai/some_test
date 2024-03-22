package __add_two_numbers

func lengthOfLongestSubstring(s string) int {
	var count, start, max int
	m := make(map[rune]int)
	for i, v := range s {
		idx, ok := m[v]
		if ok {
			count = i - idx
			for ; start <= idx; start++ {
				delete(m, rune(s[start]))
			}
		} else {
			count++
		}
		m[v] = i
		if max < count {
			max = count
		}
	}
	return max
}
