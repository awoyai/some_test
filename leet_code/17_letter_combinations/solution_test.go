package lettercombinations

import "testing"

func Test_lettercombinations(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		// {
		// 	name: "test",
		// 	args: "22",
		// 	want: []string{
		// 		"aa",
		// 		"ab",
		// 		"ac",
		// 		"ba",
		// 		"bb",
		// 		"bc",
		// 		"ca",
		// 		"cb",
		// 		"cc",
		// 	},
		// },
		{
			name: "test2",
			args: "2",
			want: []string{
				"a",
				"b",
				"c",
				"d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args); len(got) != len(tt.want) {
				t.Errorf("lettercombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
