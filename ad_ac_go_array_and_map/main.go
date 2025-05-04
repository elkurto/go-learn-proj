package main

import (
  "ad_ac_go_array_and_map/person"
  "fmt"
)

func main() {

  // ex 1 : create and print an array of :struct:Person
  ex01CreateAndPrintAnArrayOfPerson()

  // ex 2 : use :fn:make and :fn:copy on an array of :struct:Person
  ex02UseMakeAndCopyOnArrayOfPerson()

}

func ex01CreateAndPrintAnArrayOfPerson() {
  fmt.Println("\nExample 01: \n  Create and print 4 instances of Person")
  var aPerson = [...]person.Person{
    {1, "Bob", -5000},
    {2, "Sally", 11111},
    {3, "June", 26456444},
    {4, "Zev", 26456444},
  }

  for i := 0; i < len(aPerson); i++ {
    fmt.Println("    ", aPerson[i].ToString())
  }

  // slice array
  fmt.Println("\n  print slice aPerson[1:2]")
  var aPersonSlice = aPerson[1:3]
  for i := 0; i < len(aPersonSlice); i++ {
    fmt.Println("    ", aPersonSlice[i].ToString())
  }
  /*
  	// OUTPUT:
  	Example 01:
  	  Create and print 4 instances of Person
  	     Person {id:1, Name:Bob}
  	     Person {id:2, Name:Sally}
  	     Person {id:3, Name:June}
  	     Person {id:4, Name:Zev}

  	  print slice aPerson[1:2]
  	     Person {id:2, Name:Sally}
  	     Person {id:3, Name:June}
  */
}

func ex02UseMakeAndCopyOnArrayOfPerson() {
	
}
