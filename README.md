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
#### Import Package
```
import "github.com/FITLOSS/GoCalSymbol"
```
### Running GoCalSymbol

First you need to import GoCalSymbol package for using GoCalSymbol, one simplest example likes the follow :

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
