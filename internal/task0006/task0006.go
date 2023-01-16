/*
https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go
*/

package task0006

type Resp [3]uint64

func Solution(prod uint64) Resp {
	var (
		resp                 = Resp{}
		current_value uint64 = 0
		x             uint64 = 0
		fib1          uint64 = 0
		fib2          uint64 = 0
	)
	for current_value < prod {
		x += 1
		fib1 = fib(x)
		fib2 = fib(x + 1)
		current_value = fib1 * fib2
	}
	if current_value == prod {
		resp[0] = fib1
		resp[1] = fib2
		resp[2] = 1
		return resp

	}

	resp[0] = fib1
	resp[1] = fib2
	resp[2] = 0
	return resp
}

func fib(n uint64) uint64 {
	var (
		a uint64 = 1
		b uint64 = 1
		i uint64 = 0
	)

	for i = 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return a
}
