package main

func mySqrt1(x int) int {
	left := 0.0
	right := float64(x)
	res := -1.0

	for ; left <= right; {
		// 整除
		mid := float64(left) + float64((right-left)/2)
		if mid*mid <= float64(x) {
			res = mid
			// ----mid*mid----x---->
			left = mid + 0.0001
		} else {
			// ---x-----mid*mid----->
			right = mid - 0.0001
		}
	}
	println(res)
	return int(res)
}

func mySqrt(x int) int {
	left := 0
	right := x
	res := 0

	for ; left <= right; {
		// 整除
		mid := left + (right-left)/2
		if mid*mid <= x {
			res = mid
			// ----mid*mid----x---->
			left = mid + 1
		} else {
			// ---x-----mid*mid----->
			right = mid - 1
		}
	}
	return res
}

func main() {
	res := mySqrt(1)
	println(res)
}
