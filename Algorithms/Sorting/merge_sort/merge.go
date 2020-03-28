package merge_sort

func Merge(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	//find the middle index and split array
	mid := len(array) / 2
	left := array[:mid]
	//right := array[mid:]
	right := array[mid:]

	//then  we sort the left and right array with recursive method
	leftsorted := Merge(left)
	rightsorted := Merge(right)
	return mergeSort(leftsorted, rightsorted)
}

func mergeSort(left, right []int) []int {
	//result := make([]int,0)
	result := make([]int, 0)
	i, j := 0, 0
	/*if*/
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//left和right这两个切片的元素数量不一定相等，所以肯定有一个没有加到result里面
	if i < len(left) {
		//result = append(result, left[i])
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		//result = append(result, right[j])
		result = append(result, right[j:]...)
	}
	return result
}
