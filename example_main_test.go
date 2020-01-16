package main

import "fmt"

var store1 = [] int{13, 14, 12, 13, 16, 16, 12}
var store2 = [] int{16, 12, 13, 14, 15, 15, 12}
var store3 = [] int{12, 12, 15, 14, 12, 16, 17}
var income1 = incomeOfWeek(store1)
var income2 = incomeOfWeek(store2)
var income3 = incomeOfWeek(store3)
func ExampleIncome1OfWeek()  {
	fmt.Println(income1)
	//Output: 96
};func ExampleIncome2OfWeek()  {
	fmt.Println(income2)
	//Output: 97
};func ExampleIncom3OfWeek()  {
	fmt.Println(income3)
	//Output: 98
}

func ExampleCompareOfWeekIncome(){
	bestIncomeOfWeek := compareOfWeekIncome(income1,income2,income3)
	fmt.Println(bestIncomeOfWeek)
	//Output: 3
}