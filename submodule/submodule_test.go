package submodule

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "positive", nums: []int{1, 2, 3}, want: 6},
		{name: "negative", nums: []int{-1, -2, -3}, want: -6},
		{name: "mixed", nums: []int{-1, 2, -3, 4}, want: 2},
		{name: "empty", nums: []int{}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.nums); got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
