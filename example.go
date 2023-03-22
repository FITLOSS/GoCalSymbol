package main

import (
  "fmt"
  "github.com/FITLOSS/GoCalSymbol"
)
func main() {
	//define a compute rule
	algo_str := "-(お*は*2+よ)^3+(う*8)^2"
	//new a reuse struct with max rule length + 1
	cal := CalSymbol.NewStruct(len(algo_str) + 1)
	//set rule
	cal.GiveRule(algo_str)
	//set value of symbol
	cal.Set('お', 770.0)
	cal.Set('は', 20)
	cal.Set('よ', 20)
	cal.Set('う', 9)
	//compute result
	result := cal.Compute()
	fmt.Println(algo_str,"=", result)
}
