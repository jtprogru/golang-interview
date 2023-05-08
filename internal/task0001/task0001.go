// Что выведет следующая программа и почему?
// Сделай так, чтобы тест не падал
package task0001

type Person struct {
	Name string
}

type Resp struct {
	First  *Person
	Second *Person
}

func changeName(person *Person) {
	*person = Person{
		Name: "Alice",
	}
}

func Solution(origin *Person) Resp {
	var resp = Resp{}
	resp.First = origin
	resp.Second = &Person{Name: "John"}

	changeName(resp.First)
	return resp
}
