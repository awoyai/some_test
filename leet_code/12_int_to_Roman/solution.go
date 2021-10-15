package _2_int_to_Roman

var (
	m = map[string][]string{
		"I": []string{"I", "V", "X"},
		"X": []string{"X", "L", "C"},
		"C": []string{"C", "D", "M"},
		"M": []string{"M", "Ⅴ", "Ⅹ"},
	}
)
func intToRoman(num int) string {
	return helper(num, "I")
}
// 递归
func helper(num int, k string) string {
	if num == 0 {
		return ""
	}
	gInt := num % 10
	i := gInt / 5
	i2 := gInt % 5
	s := helper2(i2 , m[k][0])
	if i >= 1 {
		if gInt == 9 {
			s = m[k][0] + m[k][2]
		}else {
			s = m[k][1] + helper2(i2 , m[k][0])
		}
	}
	if gInt == 4 {
		s = m[k][0] + m[k][1]
	}
	return helper(num/10, m[k][2]) + s
}

func helper2(num int, s string) string {
	ss := ""
	for i := 0; i < num; i++ {
		ss+=s
	}
	return ss
}