// модуль 23 урок 5 несколько возвращаемых значений из функции
package main

import (
	"fmt"
	"log"
)

const size = 10
const newSize = 1

func main() {
	data := [size]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(data)

	newData := [newSize]int{1}
	//как из одной функции получить две разные переменные?:
	min, max := minAndMax(data)
	fmt.Println(min)
	fmt.Println(max)

	a, _ := ignoreSecond() //как вывести только одну переменную, из функции возврающей две
	fmt.Println(a)         //_ позволяет проигнорировать одну из переменных

	sum, err := calculateWithError(newData) //так обрабатываются обычно ошибки в go
	if err != nil {
		log.Fatal("calculate error: %v\n", err)
	}
	fmt.Println(sum)

	s, d := sumAndDiff(10, 5)
	fmt.Println(s, d)
}

// напишем функцию возвращающую две переменные
func minAndMax(input [size]int) (int, int) {
	if size == 0 {
		return -1, -1
	}
	max := input[0]
	min := input[0]

	for i := 0; i < size; i++ {
		if input[i] > max {
			max = input[i]
		}
		if input[i] < min {
			min = input[i]
		}
	}
	return min, max
}

// напишем еще одну функцию
func ignoreSecond() (int, int) {
	return 3, 7

}

// функции возвращающие две переменные очень часто встречаются в go в случае когда
// возвращается также ошибка вместе с переменной, и ошибку или переменную можно игнорировать
// продемонстрируем на следующем примере
func calculateWithError(input [newSize]int) (int, error) {
	if newSize == 0 {
		return -1, fmt.Errorf("empty size of input")
	}
	sum := 0
	for i := 0; i < newSize; i++ {
		sum += input[i]
	}
	return sum, nil //в случае когда ошибок не возникло возвращаем nil
}

// расмотрим несколько возвращаемые именованных значений
func sumAndDiff(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}
