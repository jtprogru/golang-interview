// Функция должна возвращать объект Resp
// В параметре A должны сожержаться все числа
// из диапазона start <= i < end
// Сделать так, чтобы тесты проходили успешно

package task0002

import (
	"sort"
	"sync"
)

type Resp struct {
	A []int
}

func Solution(s int, e int) *Resp {
	var res = &Resp{}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := s; i < e; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			res.A = append(res.A, num)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	sort.Ints(res.A)
	return res
}
