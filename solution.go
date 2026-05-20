package main

func twoSum(nums []int, target int) []int {
	var indexList []int

	for ind, v := range nums {
		//fmt.Println("V: ", v)
		//fmt.Println("ind: ", ind)
		for ind2, v2 := range nums {
			//fmt.Println("V2: ", v2)
			//fmt.Println("ind2: ", ind2)

			if ind == ind2 {
				continue
			}

			sum := v + v2
			//fmt.Println("Sum: ", sum)

			if sum == target {
				indexList = append(indexList, ind, ind2)

				return indexList
			}
		}
	}

	return nil
}
