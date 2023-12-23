package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseInputString(inputString string) (bool, string, string, string, error) {
	reArabic := regexp.MustCompile(`^\s*(\d+)\s*([-+*/])\s*(\d+)\s*$`)
	matchArabic := reArabic.FindStringSubmatch(inputString)
	isArabic := true

	if matchArabic != nil {
		return isArabic, matchArabic[1], matchArabic[2], matchArabic[3], nil
	}

	reRoman := regexp.MustCompile(`^\s*([IVX]+)\s*([-+*/])\s*([IVX]+)\s*$`)
	matchRoman := reRoman.FindStringSubmatch(inputString)

	if matchRoman != nil {
		isArabic = false
		return isArabic, matchRoman[1], matchRoman[2], matchRoman[3], nil
	}

	return false, "", "", "", fmt.Errorf("введенные данные не соответствуют ожидаемому формату")
}

func toArabic(roman string) int {
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

func toRoman(arabic int) string {
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

func isPositive(number int) bool {
	return number >= 1
}

func arithmeticOperation(num1, num2 int, op string) int {
	var result int

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	return result
}

func main() {
	fmt.Print("Введите строку: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString := scanner.Text()

		isArabic, num1, operator, num2, err := parseInputString(inputString)

		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		if !isInRange(toArabic(num1)) || !isInRange(toArabic(num2)) {
			fmt.Println("Ошибка: число выходит за пределы допустимого диапазона")
			return
		}

		result := arithmeticOperation(toArabic(num1), toArabic(num2), operator)

		if !isArabic {
			if isPositive(result) {
				fmt.Printf("Результат: %v\n", toRoman(result))
				return
			} else {
				fmt.Println("Ошибка: результат меньше I")
				return
			}
		}

		fmt.Printf("Результат: %v\n", result)
	}
}
