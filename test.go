package main

import (
		"fmt"
)

func main() {
	//var age int = 1;
	//aa := 2;
	//fmt.Println(aa)
	//iftest()
	//fortest()
	//fmt.Println(funcTestMax(1,2))
	//a,b := funcTestMutRet(1,"2")
	//fmt.Println(a,b)
	//
	//var c1 Circle
	//c1.radius = 10.00
	//fmt.Println(c1.getArea())
	//fmt.Println(testDefer1())

	//fmt.Println(Fa())
	//fmt.Println(testDefer2())

	//fmt.Println(testDefer3())
	fmt.Println(testDefer4)
}

func testDefer1() int {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func testDefer2() (x int) {
	defer func() {
		x += 1
	}()
	return 5
}

func testDefer3() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func testDefer4(x int)  {
	defer func(x int) {
		x += 1
	}(x)
	return 5
}

func Fa()(result int)  {
	return
}
//if语句
func iftest()  {
	bb := 3
	if bb == 2 {
		fmt.Println(2)
	} else {
		fmt.Println(3);
	}

	fmt.Println(bb)
}

func fortest() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++;
	}

	for j := 7;j <= 9;j++ {
		fmt.Println(j)
	}

	//for  {
	//	fmt.Println("死循环")
	//}
}

func funcTestMax(num1, num2 int) int {
	var res int
	if num1 > num2{
		res = num1
	} else {
		res = num2
	}

	return res
}

func funcTestMutRet(str1 int,str2 string) (int, string) {
	return str1,str2
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}




