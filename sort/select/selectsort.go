package _select

func SelectSort(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 0; i < n-1; i++ {
		// 每次从第 i 位开始，找到最小的元素
		min := list[i] // 最小数
		minIndex := i  // 最小数的下标
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				// 如果找到的数比上次的还小，那么最小的数变为它
				min = list[j]
				minIndex = j
			}
		}
		// 这一轮找到的最小数的下标不等于最开始的下标，交换元素
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

// SelectGoodSort 算法改进
func SelectGoodSort(list []int) {
	n := len(list)
	// 只需循环一半
	for i := 0; i < n/2; i++ {
		minIndex := i // 最小值下标
		maxIndex := i // 最大值下标
		// 在这一轮迭代中要找到最大值和最小值的下标
		for j := i + 1; j < n-i; j++ {
			// 找到最大值下标
			if list[j] > list[maxIndex] {
				maxIndex = j // 这一轮这个是大的，直接 continue
				continue
			}
			// 找到最小值下标
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if maxIndex == i && minIndex != n-i-1 {
			// 如果最大值是开头的元素，而最小值不是最尾的元素
			// 先将最大值和最尾的元素交换
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			// 然后最小的元素放在最开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			// 如果最大值在开头，最小值在结尾，直接交换
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			// 否则先将最小值放在开头，再将最大值放在结尾
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}
	}
}