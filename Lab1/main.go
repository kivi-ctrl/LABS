package main

import (
	"fmt"
	"lab1/funct"
	"time"
)

func main() {
	today := time.Now() //1
	fmt.Println(today)
	A := int(5) //2
	B := float64(2.4548)
	S := string("Hello world!")
	C := bool(true)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(S)
	fmt.Println(C)
	A1 := 15 //3
	fmt.Println(A1)
	B1 := 1.4512
	fmt.Println(B1)
	var a2 int64 //4
	var b2 int64
	var k string
	fmt.Println("Введите первое целое число:")
	fmt.Scan(&a2)
	fmt.Println("Выберете действие(+;-;*;/;%):")
	fmt.Scan(&k)
	fmt.Println("Введите второе целое число:")
	fmt.Scan(&b2)
	switch {
	case k == "+":
		fmt.Println("Сумма чисел равна " + fmt.Sprint(a2+b2))
	case k == "-":
		fmt.Println("Разница чисел равна " + fmt.Sprint(a2-b2))
	case k == "*":
		fmt.Println("Произведение чисел равно " + fmt.Sprint(a2*b2))
	case k == "/":
		fmt.Println("Деление чисел равно " + fmt.Sprint(a2/b2))
	case k == "%":
		fmt.Println("Остаток от деления равен " + fmt.Sprint(a2%b2))
	}
	var f1 float32 //5
	var f2 float32
	fmt.Println("Введите первое число:")
	fmt.Scan(&f1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&f2)
	fmt.Println("Сумма равна:" + fmt.Sprint(funct.Sum(f1, f2)))
	fmt.Println("Разница равна:" + fmt.Sprint(funct.Sub(f1, f2)))
	var l float32 //6
	var m float32
	var h float32
	fmt.Println("Введите первое число:")
	fmt.Scan(&l)
	fmt.Println("Введите второе число:")
	fmt.Scan(&m)
	fmt.Println("Введите третье число:")
	fmt.Scan(&h)
	fmt.Println("Среднее арифметическое равно " + fmt.Sprint((l+m+h)/3))
}
