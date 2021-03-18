package main

func firstUniqChar2(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar(s string) int {
	// 正确的处理方式是将 string 转为 rune 数组
	runeArr := []rune(s)
	hash := make(map[string]int)
	for i := 0; i < len(runeArr); i++ {
		cur := string(runeArr[i])
		hash[cur] = hash[cur] + 1
	}
	println(hash)
	for i := 0; i < len(runeArr); i++ {
		cur := string(runeArr[i])
		if hash[cur] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	res := firstUniqChar("leetcode")
	println(res)
	res1 := firstUniqChar("哈哈哈哈哈哈哈哈呵")
	println(res1)
}
