package main

import (
			"fmt"
			"math/rand"
			"time"
			"os"
) 

func main(){
//изменения внесены
cout := 0
guess := -1
var c string
var n string
var one int = 0
var thy int = 0
var f int
var nn int
fmt.Println();

fmt.Print("Выбери способы операции вычисления: ")	//вывести сообщение
fmt.Fscan(os.Stdin, &c)								//ввод с клавиатуры
//fmt.Println(c)
fmt.Println();
//выбор вычисления

switch c {
case "+":
	fmt.Println("ты выбрал сложение")
	n = "one"	
case "-":
	fmt.Println("ты выбрал вычитание")
	n = "thy"	
case "*":
	fmt.Println("ты выбрал умножение")
	n = "three"	
case "/":
	fmt.Println("ты выбрал деление")
	n = "four"	
default:
	fmt.Println("Извини, такого нет :C ")
}

fmt.Println()

for cout < 10 && guess != one && guess != thy {

	rand.Seed(time.Now().UnixNano())
	one = rand.Intn(101)
	//fmt.Println(one)
	rand.Seed(time.Now().UnixNano())
	thy = rand.Intn(200)
	//fmt.Println(thy)

	fmt.Println("Осталось ",10 - cout," попыток")
	
	 if n == "one"{ 
	 	fmt.Print(" ", one ," + ", thy, " = ")
	 	fmt.Fscan(os.Stdin, &f)	
		b := one + thy
			if b == f {
				nn = nn + 1
			}
	}
	if n == "thy"{  
		fmt.Print(" ", one ," - ", thy, " = ")
	 	fmt.Fscan(os.Stdin, &f)	
		b := one - thy
			if b == f {
				nn = nn + 1
			}
	}
	if n == "three"{  
		fmt.Print(" ", one ," * ", thy, " = ")
	 	fmt.Fscan(os.Stdin, &f)	
		b := one * thy
			if b == f {
				nn = nn + 1
			}
	}
	if n == "four"{  
		fmt.Print(" ", one ," / ", thy, " = ")
	 	fmt.Fscan(os.Stdin, &f)	
		b := one / thy
			if b == f {
				nn = nn + 1
			}
	}
	cout++

}

fmt.Println("правильно: ", nn)
}
