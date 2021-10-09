package _12_has_path_sum

import "testing"

func BenchmarkHasPathSum(b *testing.B) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root:      &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				targetSum: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := HasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				b.Errorf("HasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
