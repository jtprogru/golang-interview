/*
Given a string `s` and an integer `k, reverse the first `k` characters
for every `2k` characters counting from the start of the string.

If there are fewer than `k` characters left, reverse all of them.
If there are less than `2k` but greater than or equal to `k` characters,
then reverse the first `k` characters and leave the other as original.
*/
package task0011

func Solution(s string, k int) string {
	bb := []byte(s)
	for i := 0; i < len(s); {
		a := i
		b := i + k - 1
		if b > (len(bb) - 1) {
			b = len(bb) - 1
		}
		for a < b {
			tmp := bb[a]
			bb[a] = bb[b]
			bb[b] = tmp
			a++
			b--
		}
		i = i + 2*k
	}

	return string(bb)
}
