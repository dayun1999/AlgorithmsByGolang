package counting_sort

//countingSort can noly sort k integers <-{0,1,2,......k-1}
//max is the maximum element in the array
//here we define arr {2,3,3,7,5,5,5,6}
func CountingSort(arr []int, max int) []int{
	//定义一个[]int来存放arr里每个元素出现的次数
	count := make([]int, max+1)
	out := make([]int,0)
	//下面开始计算元素出现的次数，并将次数与元素对应的索引存起来
	for _, item := range arr{
		count[item]++
	}
	//上面操作完成之后,count切片应该是这样的{0,0,1,2,0,3,0,1,1}
	for i,item := range count{
		for j:=0;j<item;j++{
			out = append(out,i)
		}
	}

	return out
}
