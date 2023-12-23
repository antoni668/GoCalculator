package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseInputString(inputString string) (string, string, string, error) {
	reArabic := regexp.MustCompile(`^\s*(\d+)\s*([-+*/])\s*(\d+)\s*$`)
	matchArabic := reArabic.FindStringSubmatch(inputString)

	if matchArabic != nil {
		return matchArabic[1], matchArabic[2], matchArabic[3], nil
	}

	reRoman := regexp.MustCompile(`^\s*([IVX]+)\s*([-+*/])\s*([IVX]+)\s*$`)
	matchRoman := reRoman.FindStringSubmatch(inputString)

	if matchRoman != nil {
		return matchRoman[1], matchRoman[2], matchRoman[3], nil
	}

	return "", "", "", fmt.Errorf("введенные данные не соответствуют ожидаемому формату")
}

func romanToArabic(roman string) int {
	regex := regexp.MustCompile("^[0-9]+$")

	if regex.MatchString(roman) {
		num, _ := strconv.Atoi(roman)
		return num
	}

	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}

	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[string(roman[i])]

		if value < prevValue {
			result -= value
		} else {
			result += value
		}

		prevValue = value
	}

	return result
}

func arabicToRoman(arabic int) string {
	arabicNumerals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(arabicNumerals); i++ {
		for arabic >= arabicNumerals[i] {
			result += romanNumerals[i]
			arabic -= arabicNumerals[i]
		}
	}

	return result
}

func isInRange(number int) bool {
	return number >= 1 && number <= 10
}

func main() {
	fmt.Print("Введите строку: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString := scanner.Text()

		num1, operator, num2, err := parseInputString(inputString)

		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		if !isInRange(romanToArabic(num1)) || !isInRange(romanToArabic(num2)) {
			fmt.Println("Ошибка: число выходит за пределы допустимого диапазона")
			return
		}

		fmt.Printf("Первое число: %s\n", num1)
		fmt.Printf("Оператор: %s\n", operator)
		fmt.Printf("Второе число: %s\n", num2)
	}
}
