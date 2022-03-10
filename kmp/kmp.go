package main

import "fmt"

func getNext(p string) (next []int) {
	plen := len(p)
	next = make([]int, plen)

	j := 0
	k := -1
	next[0] = -1
	for j < plen-1 {
		a := p[j]
		var b uint8
		if k > -1 {
			b = p[k]
		}

		//这里,k表示next[j],且s[k]表示前缀,s[j]表示后缀
		//注:k==-1表示未找到k前缀与k后缀相等,首次分析可先忽略
		if k == -1 || a == b { //"abcabb" //j相等位置，就是下一位不等时可移动位
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}

		fmt.Printf("j:%s,k:%s,next[%d]==%d \n", string(a), string(b), j, next[j])
	}
	return
}

//优化版
func getNex2(p string) (next map[int]int) {
	next = make(map[int]int)

	j := 0
	k := -1
	next[0] = -1
	plen := len(p)
	for j < plen-1 {
		a := p[j]
		var b uint8
		if k > -1 {
			b = p[k]
		}

		if k == -1 || a == b { //"abcabb" //比较位j前字符串最在重复了串
			j++
			k++
			if p[j] == p[k] {
				next[j] = next[k]
			} else {
				next[j] = k
			}
		} else {
			k = next[k]
		}

		fmt.Printf("j:%s,k:%s,next[%d]==%d \n", a, b, j, next[j])
	}
	return
}

func kmp(t, p string) int {
	i := 0 // 主串的位置
	j := 0 // 模式串的位置
	next := getNext(p)
	fmt.Println(next)
	tlen := len(t)
	plen := len(p)
	for i < tlen && j < plen {
		if j == -1 || t[i] == p[j] {
			i++
			j++
		} else {
			// i不需要回溯了
			// i = i - j + 1

			j = next[j]
		}
	}

	if j == plen {
		return i - j
	} else {
		return -1
	}
}

func main() {
	t := "gabfabacabb"
	p := "abac"

	index := kmp(t, p)
	fmt.Println(index)

}
