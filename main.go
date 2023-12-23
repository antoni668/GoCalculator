package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

		fmt.Printf("Первое число: %s\n", num1)
		fmt.Printf("Оператор: %s\n", operator)
		fmt.Printf("Второе число: %s\n", num2)
	}
}
