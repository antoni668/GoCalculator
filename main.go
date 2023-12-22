package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите строку: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString := scanner.Text()
		fmt.Printf("Вы ввели: %s\n", inputString)
	}
}
