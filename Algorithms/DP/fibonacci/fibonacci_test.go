package fibonacci

import "testing"

//这里用uint64,防止后面计算10000的时候为负数
var addCases = []uint64{5,10,20,50,100,1000,10000}

func TestFibonacci(t *testing.T) {
	for _, tc := range addCases{
		result := Fibonacci(tc)
		t.Logf("斐波那契数列的第%d个值为%d\n",tc,result)
	}
}
