package __add_two_numbers

import (
	"reflect"
	"testing"
)

func BenchmarkAddTwoNumbers(b *testing.B) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{
								Val:  9,
								Next: &ListNode{
									Val:  9,
									Next: &ListNode{
										Val:  9,
										Next: &ListNode{
											Val:  9,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val:  9,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{
							Val:  9,
							Next: &ListNode{
								Val:  0,
								Next: &ListNode{
									Val:  0,
									Next: &ListNode{
										Val:  0,
										Next: &ListNode{
											Val:  1,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				b.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
