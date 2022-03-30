package main

import "fmt"

func main(){
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9()
	task10()
	task11()
	task12()

}

func task1(){
	var number int
	fmt.Scan(&number)
	number1, number2, number3 := number/100, (number%100)/10, (number%100)%10
	println("Сумма:= ", number1+number2+number3, "; Произведение: ", number1*number2*number3)
}

func task2(){
	var number int
	fmt.Scan(&number)
	finalNumber := (number%100)/10
	println(finalNumber)
}

func task3(){
	var number int
	fmt.Scan(&number)
	finalNumber := (number%1000)/100
	println(finalNumber)
}

func task4(){
	var number int
	fmt.Scan(&number)
	finalNumber := (number%10000)/1000
	println(finalNumber)
}

func task5(){
	var a1, d int
	var element = 30
	fmt.Scan(&a1, &d)
	println(a1 + d*(element-1))
}

func task6(){
	var a1, d, n int
	fmt.Scan(&a1, &d, &n)
	println((a1 + (a1 + d*(n-1))) * n/2)
}

func task7(){
	var chislo, n int
	fmt.Scan(&chislo)
	n = chislo%7
	println(n)
}

func task8(){
	var n, x int
	fmt.Scan(&n)
	x = (n%12)+1
	println(x)
}

func task9(){
	var n int
	fmt.Println("Введите величину временного интервала (в минутах)")
	fmt.Scan(&n)
	println(n, " мин - это", n/60, " час ", n%60, "мин")
}

func task10(){
	var (
		n int
		z [3]int
	)
	fmt.Println("Введите 3х-значное число")
	fmt.Scan(&n)
	z[0], z[1], z[2] = n/100, (n%100)/10, (n%100)%10
	n = z[1]*100 + z[0] * 10 + z[2]
	println(n)
}

func task11(){
	var (
		n, z int
	)
	fmt.Println("Введите 3х-значное число")
	fmt.Scan(&n)
	z = (n%100%10)*1000 + n
	println(z)
}

func task12(){
	var (
		n int
		z [4]int
	)
	fmt.Println("Введите 4х-значное число")
	fmt.Scan(&n)
	z[0], z[1], z[2], z[3] = n/1000, (n%1000)/100, (n%1000%100)/10, n%1000%100%10
	n = z[2]*1000 + z[3]*100 + z[0] * 10 + z[1]
	println(n)
}



