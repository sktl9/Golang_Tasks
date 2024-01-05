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

//9(5kuy)

func MoveZeros(arr []int) []int {
    nonZeroIndex := 0

    for i, num := range arr {
        if num != 0 {
            arr[i], arr[nonZeroIndex] = arr[nonZeroIndex], arr[i]
            nonZeroIndex++
        }
    }

    return arr
}

//10(6kuy)
func ArrayDiff(a, b []int) []int {
	result := make([]int, 0)

   bOccurrences := make(map[int]int)
   for _, num := range b {
	   bOccurrences[num]++
   }

   for _, num := range a {
	   if bOccurrences[num] == 0 {
		   result = append(result, num)
	   }
   }

   return result
}

//11(6kuy) (js)
function solution(words, millipede = "") {
	if (!words.length) return true;
	const results = [];
	for(let i = 0; i < words.length; i ++){
	  const word = words[i];
	  const slicedWords = words.slice(0, i).concat(words.slice(i+1))
	  if (millipede === "") {
		results.push(solution(slicedWords, word))
	  }
	  if (word.slice(-1) === millipede[0]) {
		results.push(solution(slicedWords, word + millipede))
	  }
	  else if (word[0] === millipede.slice(-1)) {
		results.push(solution(slicedWords, millipede + word))
	  }
	  else {
		results.push(false)
	  }
	}
	return results.filter(result => result).length > 0;
  }

  //12(6kuy)
  func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

//13(6kuy)
func QueueTime(customers []int, n int) int {
	if n <= 0 || len(customers) == 0 {
			return 0
		}
	
		tills := make([]int, n)
	
		for _, time := range customers {
	
			minIndex := findMinIndex(tills)
	
			tills[minIndex] += time
		}
	
		return findMax(tills)
	}
	
	func findMinIndex(arr []int) int {
		minIndex := 0
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			}
		}
		return minIndex
	}
	
	func findMax(arr []int) int {
		max := arr[0]
		for _, val := range arr {
			if val > max {
				max = val
			}
		}
		return max
	}

//14(5kuy)(js)
function pigIt(str){
	const words = str.split(" ");
 
   for (let i = 0; i < words.length; i++) {
	 const word = words[i];
 
	 if (/^[a-zA-Z]+$/.test(word)) {
	   words[i] = word.slice(1) + word[0] + "ay";
	 }
   }
 
   return words.join(" ");
 }

 //15 (6kuy)
 func DigitalRoot(n int) int {
    if n < 10 {
        return n
    }
    return DigitalRoot(sumDigits(n))
}

func sumDigits(n int) int {
    if n == 0 {
        return 0
    }
    return n%10 + sumDigits(n/10)
}
//16 (5kuy)
func RGB(r, g, b int) string {

	r = clamp(r)
	g = clamp(g)
	b = clamp(b)


	result := fmt.Sprintf("%02X%02X%02X", r, g, b)
	return result
}

func clamp(value int) int {
	return int(math.Max(0, math.Min(float64(value), 255)))
}