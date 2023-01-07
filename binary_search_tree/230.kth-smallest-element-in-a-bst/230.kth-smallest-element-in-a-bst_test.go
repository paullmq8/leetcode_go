package binary_search_tree

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  1,
						Left: nil,
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				k: 1,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				root: &TreeNode{
					Val:   5,
					Left:  &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 1,
								Left: nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val: 4,
							Left: nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: nil,
						Right: nil,
					},
				},
				k: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
