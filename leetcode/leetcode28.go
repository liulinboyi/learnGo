package main

func getNext(needle string) []int {
	next := make([]int, len(needle), len(needle))
	next[0] = 0
	j := 0
	for i := 1; i < len(needle); i++ {
		for ; j >= 1 && needle[i] != needle[j]; {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	return next
}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := getNext(needle)
	//println(next, "next")

	i := 0
	j := 0

	for ; i < len(haystack); {
		for ; j >= 1 && haystack[i] != needle[j]; {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
		i++
	}

	return -1
}

func main() {
	res := strStr("hello", "ll")
	println(res)
}
