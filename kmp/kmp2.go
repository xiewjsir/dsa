package main

import "fmt"

//暴力匹配
func ViolenceMatch(s, p string) int {
	i := 0 // 主串的位置
	j := 0 // 模式串的位置
	for i < len(s) && j < len(p) {
		if s[i] == p[j] { // 当两个字符相同，就比较下一个
			i++
			j++
		} else {
			i = i - j + 1 // 一旦不匹配，i后退
			j = 0         // j归0
		}
	}
	if j == len(p) {
		return i - j
	} else {
		return -1
	}
}

func GetNext(s string) []int {
	next := make([]int, len(s)+1)
	next[0] = -1
	k := -1
	j := 0
	for j < len(s)-1 {
		//fmt.Println("k=",k,"j=",j, "next =",next)
		//这里,k表示next[j-1],且s[k]表示前缀,s[j]表示后缀
		//注:k==-1表示未找到k前缀与k后缀相等,首次分析可先忽略
		if k == -1 || s[j] == s[k] {
			j++
			k++
			next[j+1] = k
		} else { //s[j]与s[k]不匹配,继续递归计算前缀s[next[k]]
			k = next[k]
		}
	}

	return next
}

func Kmp(s, p string, next []int) int {
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(p) {
		return i - j
	}
	return -1
}

func main() {
	nexts := GetNext("abac")
	index := Kmp("gabfabacabb", "abac", nexts)
	fmt.Println(index)
	fmt.Println(nexts)
}
