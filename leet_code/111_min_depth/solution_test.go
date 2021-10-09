package _11_min_depth

import "testing"

func BenchmarkMinDepth(b *testing.B) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试用例:[3,9,20,null,null,15,7]",
			args: args{root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			}},
			want: 2,
		},
	}
		for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := MinDepth(tt.args.root); got != tt.want {
				b.Errorf("MinDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
