package merge_sort

import "testing"

func TestMerge(t *testing.T) {
	for _, tc := range addCases{
		result := Merge(tc.array)
		t.Log("排序之后的数组是：",result)
	}
}
