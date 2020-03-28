package fibonacci

//Fibonacci函数使用的是bottom-up，也就是自下向上的方法来计算第n个值,空间复杂度为O(1)
func Fibonacci(n uint64) uint64{
	if n<= 1{
		return 1
	}
	var previous uint64  = 0
	var current uint64 = 1
	var result uint64 = 0
	var i uint64 = 0
	for i = 0; i < n-1;i++{
		result = previous + current
		previous = current
		current = result
	}
	return result
}
