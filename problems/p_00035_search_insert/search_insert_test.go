package p_00035_search_insert

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"case2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"case3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"case4", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
