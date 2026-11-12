package main

import (
	"fmt"

	"https/github.com/Seud-Behru/addis-eats-bacend/internal/order"
)

func getDeliveryFee(area string) (int, error) {
	fees := map[string]int{
		"Bole": 3000, "Kazanchis": 3500,
		"Piassa": 4000, "Megenagna": 4500,
	}

	fee, ok := fees[area] // comma-ok: 0, false
	if !ok {
		return 0, fmt.Errorf("unserved area: %s", area)
	}

	delete(fees, "Piassa")
	fmt.Println(len(fees))

	return fee, nil
}
type Dish struct {
    Name        string
    PriceSantim int
    IsFasting   bool
	}
type Order struct {
    Items []Dish
    Area  string
    Paid  bool
	}

func markPaidCopy(o Order)  { o.Paid = true }   // no effect
func markPaid(o *Order)     { o.Paid = true }   // works

func main (){

	fmt.Println("hellow world")

	/*var name string = "seud"
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
	fmt.Println(*y)*/

	//DAy 55 Lab 1 

	subTotal:=order.LineTotal(300, 2)
	fmt.Println("subTotal", subTotal)
	serviceFee := order.ServiceFeeSantim(subTotal)
	fmt.Println("serviceFee", serviceFee)
	deliveryFee, err := order.DeliveryFeeSantim("Bole", )
	fmt.Println("deliveryFee", deliveryFee)
	test, err := order.DeliveryFeeSantim("Megenagna", )
	fmt.Println("test", test, err)

	fmt.Println("========================================")
 // map and slice
	menu := make([]string, 0, 2)
	fmt.Println(len(menu), cap(menu))   // 0 2
	
	menu = append(menu, "Doro Wat")     // 1 2
	fmt.Println(menu)
	menu = append(menu, "Shiro")        // 2 2
	fmt.Println(menu)
	menu = append(menu, "Tibs")         // 3 4 <- reallocated
	fmt.Println(menu)
	
	// append RETURNS the slice; it does not modify in place
	//append(menu, "Kitfo")               // compile error: unused
	menu = append(menu, "Kitfo")      // correct
	fmt.Println(menu)

	more :=[]string{"mango", "papaya"}
	
	menu = append(menu, more...)   
	fmt.Println(menu)     // append another slice

	fmt.Println("========Re-slicing============")
		
	prices := []int{18500, 12000, 24000, 4500}

	newP := make([]int, len(prices))
	copy(newP, prices)
	fmt.Println("newP",newP)
 
	first2 := newP[0:2] 
	fmt.Println(first2)     // [18500 12000]
	last2  := prices[2:]  
	fmt.Println(last2)       // [24000 4500]
	all := prices[:]
	fmt.Println(all)
	first2[0 ] = 100
	fmt.Println("price[0]",prices[0])
	fmt.Println("newP[0]", newP[0])

	fmt.Println("===========map -lookaps================")

	fees := map[string]int{
    "Bole": 3000, "Kazanchis": 3500,
    "Piassa": 4000, "Megenagna": 4500,
	}
	
	fee := fees["Ayat"]           // 0 — zero value, not an error

	fmt.Println(fee)
	
    fe, ok := getDeliveryFee("Ayat")
    if ok != nil{
       fmt.Println(ok)
	}
	fmt.Println(fe)
	
	
	delete(fees, "Piassa")
	fmt.Println(len(fees))

	//structs, Pointers & Methods
	
	
	// field names: clear, and safe to reorder
	d := Dish{Name: "Shiro", PriceSantim: 12000, IsFasting: true}

	fmt.Println(d.Name)

	

	order := Order{Area: "Bole"}
 
	markPaidCopy(order)
	fmt.Println(order.Paid)     // false
	
	markPaid(&order)            // & takes the address
	fmt.Println(order.Paid)     // true
	
	p := &order
	fmt.Println(p.Area)     
}