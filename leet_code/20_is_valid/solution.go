package _0_is_valid

func IsValid(s string) bool {
	var stack []string
	m := map[string]string{"}": "{", "]": "[", ")": "("}
	for _, v := range s {
		c := string(v)
		len := len(stack)
		switch c {
		case "{", "[", "(":
			stack = append(stack, c)
		case "}", "]", ")":
			if len == 0 {
				return false
			}
			if stack[len-1] != m[c] {
				return false
			}
			stack = stack[:len-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
