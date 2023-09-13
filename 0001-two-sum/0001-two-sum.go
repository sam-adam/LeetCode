func twoSum(nums []int, target int) []int {
    var res []int

    found := false

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                found = true
                res = append(res, i, j)

                break
            }
        }

        if found == true {
            break
        }
    }

    return res
}