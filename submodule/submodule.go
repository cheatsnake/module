// Package submodule provides some simple functions
package submodule

// Sum returns the sum of an array of integers
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
