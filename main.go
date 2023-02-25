package main

import (
	"errors"
	"fmt"
	"strconv"
)

var arabic bool

func arabicToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	res := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			res = res + conversion.digit
			number -= conversion.value
		}
	}

	return res
}

func main() {

	var a, b, c, d string

	fmt.Scanln(&a, &c, &b, &d)

	if d != "" {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	v1, v2, err := values(a, b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if v1 >= 0 && v1 <= 10 && v2 >= 0 && v2 <= 10 {
		res, err := do(v1, v2, c)
		if err != nil {
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		}
		if !arabic {
			rom := arabicToRoman(res)
			fmt.Println(rom)
			return
		}
		fmt.Println(res)
	} else {
		fmt.Println("Не попал!")
	}
}

func values(a, b string) (int, int, error) {
	v1 := romanToInt(a)
	v2 := romanToInt(b)

	if (v1 == 0 && v2 != 0) || (v2 == 0 && v1 != 0) {
		return 0, 0, errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
	if v1 == 0 && v2 == 0 {
		v1, err := strconv.Atoi(a)
		if err != nil {
			return 0, 0, errors.New("Введенный текст не является числом.")
		}
		v2, err := strconv.Atoi(b)
		if err != nil {
			return 0, 0, errors.New("Введенный текст не является числом.")
		}
		arabic = true
		return v1, v2, nil
	}
	return v1, v2, nil
}

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func do(a, b int, c string) (int, error) {
	switch c {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("неверные условия")
	}
}
