/*
https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go
*/

package task0006

type Resp [3]uint64

func Solution(prod uint64) Resp {
	var f1, f2 uint64 = 0, 1
	for f1*f2 < prod {
		f1, f2 = f2, f1+f2
	}
	if f1*f2 == prod {
		return [3]uint64{f1, f2, 1}
	}
	return [3]uint64{f1, f2, 0}
}
