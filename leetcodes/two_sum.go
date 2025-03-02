package leetcodes

type TwoSum struct{}

// 37 ms

func (l *TwoSum) Solve(nums []int, target int) []int {
	var aIdx int
	var bIdx int

	for i, _ := range nums {
		for j, _ := range nums[i+1:] {
			if nums[i]+nums[j] == target {
				aIdx = i
				bIdx = j + i + 1
			}
		}
	}

	return []int{aIdx, bIdx}
}
