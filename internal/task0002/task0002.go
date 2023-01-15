// Функция должна возвращать объект Resp
// В параметре A должны сожержаться все числа
// из диапазона start <= i < end
// Сделать так, чтобы тесты проходили успешно

package task0002

type Resp struct {
	A []int
}

func Solution(s int, e int) *Resp {
	var res = &Resp{}
	for i := s; i < e; i++ {
		go func(num int, r *Resp) {
			res.A = append(r.A, num)
		}()
	}
	return res
}
