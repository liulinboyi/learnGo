package main

import (
	"fmt"
	"regexp"
)

func match(s string) bool {
	match, err := regexp.MatchString("[\\s\\S]", s)
	fmt.Println("Match: ", match, " Error: ", err)
	if err != nil {
		panic(err)
	}
	return match
}

func findString(s string) string {
	compile, _ := regexp.Compile("{{([\\s\\S]+)}}")
	return compile.FindString(s)
}

func findStringSubmatch(s string) []string {
	compile, _ := regexp.Compile("{{([\\s\\S]+)}}")
	return compile.FindStringSubmatch(s)
}

func main() {
	res := match("hello")
	fmt.Println(res)
	res1 := findString("{{hello}}")
	fmt.Println(res1)
	res2 := findStringSubmatch("{{hello}}")
	fmt.Println(res2)
}
