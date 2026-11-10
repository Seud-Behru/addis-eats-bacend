package main

import (
	"fmt"

	"https/github.com/Seud-Behru/addis-eats-bacend/internal/order"
)



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

	var  x = 100
	var y = &x
	fmt.Println(*y)

	//DAy 55 Lab 1 

	subTotal:=order.LineTotal(300, 2)
	fmt.Println("subTotal", subTotal)
	serviceFee := order.ServiceFeeSantim(subTotal)
	fmt.Println("serviceFee", serviceFee)
	deliveryFee, err := order.DeliveryFeeSantim("Bole", )
	fmt.Println("deliveryFee", deliveryFee)
	test, err := order.DeliveryFeeSantim("Megenagna", )
	fmt.Println("test", test, err)

    
	
}