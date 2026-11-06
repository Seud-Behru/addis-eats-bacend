package main

import "fmt"

func main (){

	fmt.Println("hellow world")

	var name string = "seud"
	 age := 20

	 var name1, name2 string = "abebe", "kebede"

	 const (
		name3 = "HERMELA"
		age2 = "22"

	 )

	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(age)
	fmt.Println(&age)
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(age2)


	for i :=0; i < len(name3); i++ {
		fmt.Println(name3[i])
	}

	
}