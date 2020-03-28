package counting_sort

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{2, 3, 7, 5, 5, 5, 9, 8, 8}
	result := CountingSort(arr, 9)
	t.Log("结果是：", result)
}
