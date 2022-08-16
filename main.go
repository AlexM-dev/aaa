package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите  данные: ")
	scanner.Scan()
	d := scanner.Text()
	fmt.Printf("Вы ввели данные: %v\n", d)
}
