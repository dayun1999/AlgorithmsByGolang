package radix_sort

import "testing"

func TestRadixSort(t *testing.T) {
	arr := []int{456, 56, 90, 802, 999, 53, 71}
	RadixSort(arr)
	t.Log("排序之后的数组是：", arr)
}
