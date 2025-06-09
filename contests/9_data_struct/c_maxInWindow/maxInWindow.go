package maxInWindow

func MaxInWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	dq := []int{}
	res := []int{}

	for i := 0; i < len(nums); i++ {
		if len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}

		for len(dq) > 0 && nums[dq[len(dq)-1]] <= nums[i] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)

		if i >= k-1 {
			res = append(res, nums[dq[0]])
		}
	}

	return res
}
