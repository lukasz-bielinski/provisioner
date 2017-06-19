package main

import "fmt"

//Ex1
//implement isEven int => bool function

//Kelvin new type Kelvin
type Kelvin int

//Celsius is new type
type Celsius int //type Celsius

//func sensor()  Kelvin {
//	???
//}

//func converter(sensor func() Kelvin) func() Celsius {
//	???
//}

//OneArgFunc type
type OneArgFunc func(int) int

//func curry(input func(arg1 int,arg2 int) int) func(int) func(int) int {
//or
//func curry(input func(arg1 int,arg2 int) int) func(int) OneArgFunc {
//	???
//}

func main() {
	fmt.Println("1.1a")
	//Exercise1
	//fmt.Println("isEven(1) must be false : ",isEven(1))
	//fmt.Println("isEven(2) must be true : ",isEven(2))
	//fmt.Println("isEven(3) must be false : ",isEven(3))
	//fmt.Println("isEven(4) must be true : ",isEven(4))

	fmt.Println("Exercise2")
	//fmt.Println("sensor must be 283  Kelvin : ",sensor())
	//celsiusSensor := converter(sensor)
	//fmt.Println("converted must be 10 Celsius : ",celsiusSensor())

	fmt.Println("Exercise3")
	//var curried=curry(func(arg1 int,arg2 int) int{return arg1+arg2})
	//fmt.Println(curried(2)(3))
}
