package main

func compareElement(array []int, index int, value int) int {
	if index < 0 || index >= len(array) {
		return -1 // Индекс вне границ // 1
	}
	element := array[index]
	if element > value {
		return 1 // Элемент больше // 2
	} else if element < value {
		return -1 // Элемент меньше // 3
	}
	return 0 // Элемент равен // 4
}

type Person struct {
	Name string
	Age  int
}

func compareAge(people []*Person, index int, value int) int {
	if index < 0 || index >= len(people) {
		return -1 // Индекс вне границ // 1
	}
	age := people[index].Age // Достаем возраст по индексу

	if age > value {
		return 1 // Возраст больше // 2
	} else if age < value {
		return -1 // Возраст меньше // 3
	}
	return 0 // Возраст равен // 4
}
