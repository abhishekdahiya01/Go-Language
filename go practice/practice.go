package main

import "fmt"

func main(){
	var age int = 23

	// fmt.Println(age)
	// fmt.Println(&age) //&age address
	// fmt.Printf("%p\n", &age ) //Printf me address nikalne ke liye %p use hota he

	//pointer

	// var P *int // integer pointer ka type

	// P = &age // P me age ka address store hoga

	// fmt.Printf("Age: %v\n", age)
	// fmt.Printf("Address: %p\n", P)
	// fmt.Printf("Pointer value: %v\n", *P)
	// fmt.Printf("Address P: %p\n", &P)

	changeAge(&age)

	fmt.Println(age)
	
}

func changeAge(a *int){
	*a = 45
}