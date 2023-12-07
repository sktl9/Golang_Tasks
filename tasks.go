package main

import (
	"math"
	"strconv"
	"strings"
)

// 1
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

// 2
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// 3
func SquareOrSquareRoot(arr []int) []int {
	result := make([]int, len(arr))

	for i, num := range arr {
		sqrt := math.Sqrt(float64(num))
		if sqrt == math.Floor(sqrt) {
			result[i] = int(sqrt)
		} else {
			result[i] = num * num
		}
	}

	return result
}

// 4
func Feast(beast string, dish string) bool {
	beastStart, beastEnd := beast[:1], beast[len(beast)-1:]
	dishStart, dishEnd := dish[:1], dish[len(dish)-1:]

	return strings.EqualFold(beastStart, dishStart) && strings.EqualFold(beastEnd, dishEnd)
}

// 5
func MakeNegative(x int) int {
	if x > 0 {
		return x * (-1)
	}
	return x
}

// 6(6kyu)
func Solution(str string) []string {
	var pairs []string

	for i := 0; i < len(str); i += 2 {

		if i+1 < len(str) {
			pairs = append(pairs, str[i:i+2])
		} else {
			pairs = append(pairs, string(str[i])+"_")
		}
	}

	return pairs

}

// 7(6kyu)
func WhatCentury(year string) string {

	num, err := strconv.Atoi(year)
	if err != nil {
		return "Ошибка при преобразовании года в число"
	}

	century := (num-1)/100 + 1

	suffix := "th"
	switch century % 100 {
	case 11, 12, 13:
		suffix = "th"
	default:
		switch century % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}

	result := strconv.Itoa(century) + suffix

	return result

}

// 8(6kyu)
func FindOddd(seq []int) int {
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
