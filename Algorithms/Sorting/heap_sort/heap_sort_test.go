package heap_sort

import "testing"

func TestHeapSort(t *testing.T) {
	//arr := []int{1,4,5,90,16,67,2}
	for _, tc := range addCases {
		HeapSort(tc.array)
		t.Log("排序后的数组是: ", tc.array)
	}

}

func BenchmarkHeapSort(b *testing.B) {
	arr := []int{89, 99, 67, 0, 1, 2, 1, 9, 4, 122, 32, 43}
	for i := 0; i < b.N; i++ {
		HeapSort(arr)
	}
}
