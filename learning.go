package main

import (
	"fmt"
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

func main() {
	string := "hellow"
	fmt.Println(string[1])

}
