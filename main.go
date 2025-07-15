// Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
// Подсказка: используйте битовые операции (|, &^).
package main

import "fmt"

func SetBitOne(position uint, num int64) int64 {
	return num | (1 << position)
}

func SetBitZero(position uint, num int64) int64 {
	return num &^ (1 << position)
}
func main() {

	var num int64
	var position uint
	var bit int

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	fmt.Print("Введите позицию бита (начиная с 0): ")
	fmt.Scan(&position)

	fmt.Print("Введите значение бита (0 или 1): ")
	fmt.Scan(&bit)

	if bit == 1 {
		num = SetBitOne(position, num)
	} else if bit == 0 {
		num = SetBitZero(position, num)
	} else {
		fmt.Println("Ошибка: бит должен быть 0 или 1")
		return
	}

	fmt.Println("Результат:", num)
}
