func threeSum(nums []int) [][]int {
	var res [][]int

	// sort the array by asc value
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// loop array from left to right
	for i := 0; i < len(nums) - 2; i++ {
		// since nums already sorted, if nums[i] > 0 then it's impossible to reach 0
		if nums[i] > 0 {
			break
		}

		// prevent duplicate by checking against previous value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i]+nums[j]+nums[k]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				for j < k && nums[j] == nums[j - 1] {
					j++
				}

				for j < k && nums[k] == nums[k + 1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return res
}
