package radix_sort

//radix sort 是用来解决counting sort中出现 k 的范围是 1-n^2的
//radix sort 基于counting sort
func RadixSort(arr []int){
	//找到arr里面最大的元素，才能决定
	max := getMax(arr)
	for exp := 1; max/exp > 0;{
		countingSort(arr,exp)
		exp *=10 //每次乘10
	}
}
//参数exp的作用是取arr里面每个元素的各位，十位，百位...所以exp的值为1,10,100...
//假设arr为{456,70,999,802,53,633}
func countingSort(arr []int, exp int) /*[]int*/{ //不需要返回值
	n := len(arr)
	out := make([]int,n)
	count := make([]int,10)
	for _, item := range arr{
		//分别取arr每个元素的各位、十位、百位数字
		index := (item / exp)%10
		count[index]++
	}
//这里以exp = 1 举例所以这时候count是{1,0,1,2,0,0,1,0,0,1}

//这一步很关键哦，理解这一步就OK了
	for i:=1;i<10;i++{ //一开始写成这样了for i:=1;i<n;i++，错误
		count[i] += count[i-1]
	}
//count{1,1,2,4,4,4,5,5,5,6}
	for i := n - 1; i >= 0; i-- {
		out[count[ (arr[i]/exp)%10 ] - 1] = arr[i]
		count[ (arr[i]/exp)%10 ]--
	}
	//最后还需要将arr修改
	for i, item := range out{
		arr[i] = item
	}

	//return out 返回out一点用也没有，因为这是一个副本，arr没有改变
}

//getMax返回数组中最大的那个数字
func getMax(arr []int) int{
	max := arr[0]
	for _, item := range arr{
		if item > max{
			max = item
		}
	}
	return max
}