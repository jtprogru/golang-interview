/*
https://leetcode.com/problems/count-ways-to-build-good-strings/description/

Ваш код кажется верным. Он использует рекурсивный подход
с сохранением результатов (мемоизацией), чтобы избежать повторных
вычислений для одинаковых подзадач.

Он начинает с нулевой длины и добавляет к этому значению `zero` и `one`,
пока не достигнет предела `high`. Если длина находится
в пределах `low` и `high`, он увеличивает счетчик `ans`.

Мемоизация используется для ускорения процесса путем сохранения ранее
вычисленных результатов в массиве `f`.
*/

package task0027

func Solution(low int, high int, zero int, one int) int {
	const mod int = 1e9 + 7
	var dfs func(i int) int
	f := make([]int, high+1)

	for i := range f {
		f[i] = -1
	}

	dfs = func(i int) int {
		if i > high {
			return 0
		}
		if f[i] != -1 {
			return f[i]
		}
		ans := 0
		if i >= low && i <= high {
			ans++
		}
		ans += dfs(i+zero) + dfs(i+one)
		ans %= mod
		f[i] = ans
		return ans
	}
	return dfs(0)
}
