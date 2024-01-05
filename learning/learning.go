package main

import (
	"fmt"
	"time"
)

func slice(str string) string {
	SliceArr := str[:2]
	return SliceArr
}

var arr = [5]string{"Aерера", "aлутон", "Aаумеа", "aакемаке", "Aрида"}

var planetsArr = [8]string{
	"Меркурий",
	"Венера",
	"Земля",
	"Марс",
	"Юпитер",
	"Сатурн",
	"Уран",
	"Нептун",
}

func parse(planetsArr [8]string) {
	for i, num := range planetsArr {
		fmt.Printf("Планета %s под номером %d\n", num, i+1)
	}

}

func ForArr(arr [5]string) [5]string {
	for i := range arr {
		if arr[i][0] == 'A' {
			arr[i] = "a" + arr[i][1:]
		} else if arr[i][0] == 'a' {
			arr[i] = "A" + arr[i][1:]
		}
	}
	return arr
}

type Point struct {
	X int
	Y int
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(p1)
}

func FindOdd(seq []int) int {
	dict := make(map[int]int)
	for _, num := range seq {
		dict[num] = dict[num] + 1
	}
	for key, value := range dict {
		if value%2 != 0 {
			return key
		}
	}
	return 0
}

type food interface {
	eat()
}

type fruit struct{}

func (f *fruit) eat() {
	fmt.Println("я ем фрукт")
}

type meet struct{}

func (m *meet) eat() {
	fmt.Println("я ем мясо")
}

func test() []int {
	arr := make([]int, 10)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	return arr
}

type animals interface {
	Walk()
	Running()
}

type tigr struct{}

func (t *tigr) Walk() {
	fmt.Println("Тигр идет")
}

func (t *tigr) Running() {
	fmt.Println("Тигр бежит")
}

func say(word string, ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println(word)
	ch <- 1
}

func main() {
	var arr = []int{1231, 23465, 35484658, 3456, 45}
	fmt.Println(sum(arr) / dist(arr))

}

func sum(arr []int) int {
	var result int = 0
	for _, num := range arr {
		result = result + num
	}
	return result
}

func dist(arr []int) int {
	var result int = 0
	for _ = range arr {
		result += 1
	}
	return result
}
