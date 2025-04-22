
### 00-go-fmt
- A simple go program that uses fmt.Printf(..), fn call , and for loop. 
- @see [./00-go-fmt/main.go](./00-go-fmt/main.go)


```bash
cd 00-go-fmt
go run .


  Output:
  Hello and welcome, gopher!
  i = 100
  i = 50
  i = 33
  i = 25
  i = 20
  
```

--
### 01-go-calc
- A simple interest calculator
- uses fmt.Scan, fmt.Printf, math.Pow
- @see [./01-go-calc/main.go](./01-go-calc/main.go)


```bash

cd 01-go-fmt
go run .

Output:
starting principle (1000 default): 2000
futureValue =3416.288917
futureRealValue =1819.946051

```

--
### 02-01-go-calc-ebt
- Given revenue, expense, taxRate, compute EBT (earnings before tax), profit, ratio(EBT/profit)
- Exemplifies os.Args, function definition, multiline strings (with backticks), strconv.parseFloat, and error handling.

```bash
cd 02-01-go-calc-ebt/
go run . 1100 100 10    

Output:
revenue  =   1100.00
expenses =    100.00
taxRate  =     10.00
ebt      =   1000.00
profit   =    900.00
ratio    =      1.11
```