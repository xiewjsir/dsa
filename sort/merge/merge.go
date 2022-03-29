package merge

// MergeSort 自顶向下归并排序，排序范围在 [begin,end) 的数组
func MergeSort(array []int, begin int, end int) {
	// 元素数量大于1时才进入递归
	if end-begin > 1 {
		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := begin + (end-begin+1)/2
		// 先将左边排序好
		MergeSort(array, begin, mid)
		// 再将右边排序好
		MergeSort(array, mid, end)
		// 两个有序数组进行合并
		merge(array, begin, mid, end)
	}
}

// MergeSort2 自底向上归并排序
func MergeSort2(array []int, begin, end int) {
	// 步数为1开始，step长度的数组表示一个有序的数组
	step := 1
	// 范围大于 step 的数组才可以进入归并
	for end-begin > step {
		// 从头到尾对数组进行归并操作
		// step << 1 = 2 * step 表示偏移到后两个有序数组将它们进行归并
		for i := begin; i < end; i += step << 1 {
			var lo = i                // 第一个有序数组的上界
			var mid = lo + step       // 第一个有序数组的下界，第二个有序数组的上界
			var hi = lo + (step << 1) // 第二个有序数组的下界
			// 不存在第二个数组，直接返回
			if mid > end {
				return
			}
			// 第二个数组长度不够
			if hi > end {
				hi = end
			}
			// 两个有序数组进行合并
			merge(array, lo, mid, hi)
		}
		// 上面的 step 长度的两个数组都归并成一个数组了，现在步长翻倍
		step <<= 1
	}
}
// 归并操作
func merge(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)
	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}
	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)
	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}
