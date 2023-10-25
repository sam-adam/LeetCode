func threeSumClosest(nums []int, target int) int {
	closestSum := 0
	shortestDistance := math.MaxInt32

	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)

	// loop array from left to right
	for i := 0; i < len(nums) - 2; i++ {
		// prevent duplicate by checking against previous value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum, distance := getSumAndDistance(nums, i, j, k, target)

			if distance == 0 {
				return sum
			}

			if distance < shortestDistance {
				closestSum = sum
				shortestDistance = distance
			}

			if sum < target {
				j += 1
			} else {
				k -= 1
			}
		}
		//for k - j > 1 {
		//	jx := j + 1
		//	kx := k - 1
		//	jSum, jDistance := getSumAndDistance(nums, i, jx, k, target)
		//	kSum, kDistance := getSumAndDistance(nums, i, j, kx, target)
		//
		//	if jDistance == 0 {
		//		return jSum
		//	} else if kDistance == 0 {
		//		return kSum
		//	}
		//
		//	if jDistance >= shortestDistance && kDistance >= shortestDistance {
		//		k = kx
		//		j = jx
		//
		//		for j < k && nums[k] == nums[k - 1] {
		//			k--
		//		}
		//
		//		for j < k && nums[j] == nums[j + 1] {
		//			j++
		//		}
		//	} else if jDistance > shortestDistance {
		//		k = kx
		//		closestSum = kSum
		//		shortestDistance = kDistance
		//
		//		for j < k && nums[k] == nums[k - 1] {
		//			k--
		//		}
		//	} else {
		//		j = jx
		//		closestSum = jSum
		//		shortestDistance = jDistance
		//
		//		for j < k && nums[j] == nums[j + 1] {
		//			j++
		//		}
		//	}
		//}
	}

	return closestSum
}


func getSumAndDistance(nums []int, i int, j int, k int, target int) (int, int) {
	sum := nums[i]+nums[j]+nums[k]
	distance := sum - target

	if distance < 0 {
		distance *= -1
	}

	return sum, distance
}