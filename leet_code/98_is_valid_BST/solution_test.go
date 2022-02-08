package _8_is_valid_BST

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				},
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
