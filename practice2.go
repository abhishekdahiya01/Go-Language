package main

import "fmt"

func main() {
	var mark int

	fmt.Println("Enter your mark: ")
	fmt.Scan(&mark)

	if mark <0 || mark > 100{
		fmt.Println("Invailid")
	}else if mark >= 90{
		fmt.Println("Grade A+")
	}else if mark >= 70{
		fmt.Println("Grade A")
	}else if mark >= 60{
		fmt.Println("Grade B")
	}else if mark >= 45{
		fmt.Println("Grade c")
	}else if mark >= 33{
		fmt.Println("Grade D")
	}else{
		fmt.Println("Fail")
	}


}