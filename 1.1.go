package main

import "fmt"

func main(){
	task13()
	task14()
	task15()
	task16()
	task18()
	task19()
	task20()
	task21()
	task22()
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

func task13(){
	var (
		x, y, z  int
	)
	fmt.Println("Введите 3х-значное число")
	fmt.Scan(&y)
	if y < 100 || y > 999 || (y%100)/10 == 0{
		return
	}

	z = y/100
	x = (y%100)*10 + z
	println(x)
}

func task14(){
	var (
		x, y  int
	)
	fmt.Println("Введите 3х-значное число")
	fmt.Scan(&y)
	if y < 100 || y > 999 || y%10 == 0{
		return
	}
	x = y%10 * 100 + y/10
	println(x)
}

func task15(){
	var (
		x  int
	)
	fmt.Println("Введите секунды")
	fmt.Scan(&x)
	x = x/3600
	println(x)
}

func task16(){
	var (
		x  int
	)
	fmt.Println("Введите секунды")
	fmt.Scan(&x)
	x = x - ((x/60)*60)
	println(x)
}

// task17 не понял как совсем без условных операторов сделать.

func task18(){
	var (
		h, m, grad  int
		// не использовал float, хотя было бы точнее
	)
	fmt.Println("Введите сначала h, затем m")
	fmt.Scan(&h, &m)

	if h < 1 || h > 23 || m < 0 || m > 59 {
		return
	}

	grad = ((h % 12) * 60 + m) / 2
	println(grad)

}

func task19(){
	var (
		y, minutes  int
	)
	fmt.Println("Введите y")
	fmt.Scan(&y)

	if y < 1 || y > 359 {
		return
	}
	minutes = y * 2
	h, m := minutes / 60, minutes % 60
	println(h, m)

}

func task20(){
	var (
		y  int
	)
	fmt.Println("Введите y")
	fmt.Scan(&y)

	if y < 1 || y > 359 {
		return
	}
	m := ((y * 2) % 60) * 6
	println( m)

}

func task21(){
	var (
		k, para  int
	)
	fmt.Println("Введите k")
	fmt.Scan(&k)

	if k < 0 || k > 180 {
		return
	}
	if k % 2 == 0 {
		para = k / 2
	} else {
		para = (k / 2) + 1
	}
	println(para)

}

func task22(){
	var (
		k,x, chislo, para  int
	)
	fmt.Println("Введите k")
	fmt.Scan(&k)

	if k < 0 || k > 180 {
		return
	}
	x = (k / 20) + 1

	if k % 2 == 0 {
		para = k / 2
		chislo = x * 10 + (para%10 - 1)
	} else {
		para = (k / 2) + 1
		chislo = x * 10 + (para%10 - 1)
	}

	println(chislo)

}

//task23, 24 не понял условие и, что от меня хотят


