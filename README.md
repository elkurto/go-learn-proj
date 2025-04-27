
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

--
### 02-02-go-profit-calc
- Another profit and EBT calculator.
- This program employs a reusable function to prompt for and to collect user input.
- Exemplifies custom exception handling.

```bash
cd 02-02-go-profit-calc
go run .
Enter Revenue:1100
Enter Expenses:100
Enter taxRate:10


revenue  =   1100.00
expenses =    100.00
taxRate  =     10.00
ebt      =   1000.00
profit   =    900.00
ratio    =      1.11
```

-- 
### 02-03-go-bank-sim
- A simple ATM simulator in golang
- Exemplifies, switch, conditionals, error handling, for (control loop).

```bash
cd 02-03-go-bank-sim
go run .
```

-- 
### ad_aa_go_pointer
- exemplifies passing a parameter by reference
- initializing pointer in go
- in/out parameter passed by ref in go

```python
cd  ad_aa_go_pointer/
go run main.go

# output :
after first call : color =blue
after second call : color =gray


```