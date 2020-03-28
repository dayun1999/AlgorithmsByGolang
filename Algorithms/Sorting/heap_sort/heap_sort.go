package heap_sort

//max_heapify 伪代码
/*
l = left(i)
r = right(i)
if (l <= heap-size(A) and A[l] > A[i])
then largest = l else largest = i
if (r <= heap-size(A) and A[r] > A[largest])
then largest = r
if largest = i
then exchange A[i] and A[largest]
Max_Heapify(A, largest)
*/

func maxHeapify(arr []int, n int, i int) {
	//largest := 0
	largest := i
	l := 2 * i
	r := 2*i + 1
	if l < n && arr[l] > arr[i] { //if l <= n && arr[l] > arr[i] 报错，数组越界
		largest = l
	} else {
		largest = i
	}
	if r < n && arr[r] > arr[largest] { //r <= n && arr[r] > arr[largest] 报错，数组越界
		largest = r
	}
	//swap the value of node i and node largest if largest != i
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, n, largest)
	}
	//
}

//build max heap 伪代码
/*
Converts A[1…n] to a max heap
Build_Max_Heap(A):
 for i=n/2 downto 1
 do Max_Heapify(A, i)
*/
func buildMaxHeap(arr []int) {
	size := len(arr)                 //the number of array
	for i := size / 2; i >= 0; i-- { //for i := size / 2 ; i>0 ;i-- 出错，少了一个=，最后一个元素没有参与排序
		maxHeapify(arr, size, i)
	}

}

//正式开始heap sort
func HeapSort(arr []int) {
	//build_max_heap
	buildMaxHeap(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		//一个接一个的，把堆顶的元素和最后一个节点的元素交换
		//交换堆顶和最后一个元素(最后一个元素的节点在变小)
		arr[0], arr[i] = arr[i], arr[0]
		//每交换完成一次之后都是检查是否还满足最大堆的属性，如果不满足，要执行heapify
		maxHeapify(arr, i, 0)
	}
}
