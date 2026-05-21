package main

func twoSum(nums []int, target int) []int {
	var indexList []int

	for ind, v := range nums {
		for ind2, v2 := range nums {
			if ind == ind2 {
				continue
			}

			if v+v2 == target {
				indexList = append(indexList, ind, ind2)
				return indexList
			}
		}
	}

	return nil
}
