# GoCalSymbol
If you need dynamic formula which come from database or somewhere out of programing   
It's what you need

* **You can using utf8(unicode) to define your formula as string**
* **Don't need compile executable file each time**
* **Reuse after load formular**

**Support Operator:** 
* **+ : addition**
* **- : subtraction,negative**
* **\* : multiplication**
* **/ : division**
* **^ : power** 
* **( : left brackets**
* **) : right brackets**

**Invalid Case:**
* **negative ^ (arbitrarily numer)** 

## Getting started
### Getting GoCalSymbol
#### Download Package
```sh
go get -u github.com/FITLOSS/GoCalSymbol
```
**Or**
```sh
go get -u github.com/FITLOSS/GoCalSymbol@v0.1.1
```
#### Import Package
```
import "github.com/FITLOSS/GoCalSymbol"
```
### Running GoCalSymbol

First you need to import GoCalSymbol package for using GoCalSymbol, one simplest example likes the follow :
#### Set Value One By One
```go
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
```
#### Set Value By Map
```go
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
	//set value of symbol using map
	cal.SetByMap(map[rune]float64{
		'お': 770.0,
		'は': 20,
		'よ': 20,
		'う': 9,
	})
	//compute result
	result := cal.Compute()
	fmt.Println(algo_str, "=", result)
}
```
#### Reuse After Set Rule
```go
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

	//first time use>>>>>>>>
	//set value of symbol using map
	cal.SetByMap(map[rune]float64{
		'お': 770.0,
		'は': 20,
		'よ': 20,
		'う': 9,
	})
	//compute result
	result := cal.Compute()
	fmt.Println(algo_str, "=", result)

	//second time use>>>>>>>>
	//set value of symbol using map
	cal.SetByMap(map[rune]float64{
		'お': 10.0,
		'は': 6,
		'よ': 2,
		'う': 7,
	})
	//compute result
	result = cal.Compute()
	fmt.Println(algo_str, "=", result)
}
```
