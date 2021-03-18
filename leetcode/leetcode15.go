package main

import (
	"encoding/json"
	"sort"
)

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) <= 2 {
		res := [][]int{}
		return res
	}
	res := make([][]int, 0, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {

		// 加速1：c为非负数，就不能满足a+b+c=0了
		if nums[i] > 0 {
			return res
		}

		// 加速2：跳过计算过的数据，同时防止结果重复
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		head := i + 1
		tail := len(nums) - 1

		for ; head < tail; {
			// nums[head] + nums[tail] + nums[i] = 0 => nums[i] = -(nums[head] + nums[tail])
			sum := -(nums[head] + nums[tail])
			if sum == nums[i] {
				li := []int{nums[i], nums[head], nums[tail]}
				res = append(res, li)
				// 加速3：跳过计算过的数据，同时防止结果重复
				for ; head < tail && nums[head] == nums[head+1]; {
					head++
				}
				for ; head < tail && nums[tail] == nums[tail-1]; {
					tail--
				}
			}
			if sum <= nums[i] {
				// -(nums[head] + nums[tail]) <= nums[i]
				// 数组已经排序，此时nums[head]和nums[tail]的和取负值 小于等于 nums[i]
				// 为了让他们相等，需要让nums[head]和nums[tail]的和取负值变大
				// 则需要nums[head]和nums[tail]的和变小，取负值才会变大
				// 所以需要tail--
				tail--
			} else {
				// -(nums[head] + nums[tail]) > nums[i]
				// 此处逻辑为以上逻辑取反
				head++
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{}
	res := threeSum(nums)
	//var a = []int{1,2,3,4}
	b, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	var result = string(b)
	println(result)
}
