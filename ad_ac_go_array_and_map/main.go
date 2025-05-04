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
  // create a new array of person.Person instances
  var aPerson = make([]person.Person, 5, 10)
  aPerson[0] = person.Person{Id: 0, Name: "Smith", BirthDate: 134203470}
  aPerson[1] = person.Person{Id: 1, Name: "Jones", BirthDate: 134203470}
  aPerson[2] = person.Person{Id: 2, Name: "Salinas", BirthDate: 134203470}
  aPerson[3] = person.Person{Id: 3, Name: "Xu", BirthDate: 134203470}
  aPerson[4] = person.Person{Id: 4, Name: "Babar", BirthDate: 134203470}

  // append to aPerson
  aPerson = append(aPerson, person.Person{5, "ONeil", 1234456400})
  aPerson = append(aPerson, person.Person{6, "Akuneli", 1234456400})
  aPerson = append(aPerson, person.Person{7, "Yamamoto", 1234456400})

  fmt.Println("\n  var aPerson = make([]person.Person, 5, 10) // then populate")
  for i := 0; i < len(aPerson); i++ {
    fmt.Println("    ", aPerson[i].ToString())
  }
  /*
  		//output
  	  var aPerson = make([]person.Person, 5, 10) // then populate
  	     Person {id:0, Name:Smith}
  	     Person {id:1, Name:Jones}
  	     Person {id:2, Name:Salinas}
  	     Person {id:3, Name:Xu}
  	     Person {id:4, Name:Babar}
  	     Person {id:5, Name:ONeil}
  	     Person {id:6, Name:Akuneli}
  	     Person {id:7, Name:Yamamoto}
  	    aPersonSliceA[0] =Person {id:1, Name:Jones}
  	    aPersonSliceA[1] =Person {id:2, Name:Salinas}

  */
  // swap aPerson[1:3] with aPerson[4:6]
  var aPersonSliceA = make([]person.Person, 2) // make a temp array
  copy(aPersonSliceA[0:2], aPerson[1:3])       // copy( dest, src )  // backwards
  for i := 0; i < len(aPersonSliceA); i++ {
    fmt.Printf("    aPersonSliceA[%d] =%s\n", i, aPersonSliceA[i].ToString())
  }
  var aPersonSliceB = aPerson[4:6]       // optional make a source slice
  copy(aPerson[1:3], aPersonSliceB)      // copy to [1:3] from [4:6]
  copy(aPerson[4:6], aPersonSliceA[0:2]) // copy to [4:6] from aPersonSliceA[0:2]

  fmt.Println("\n  aPerson after swapping [1:3] and [4:6]")
  for i := 0; i < len(aPerson); i++ {
    fmt.Println("    ", aPerson[i].ToString())
  }
  /* // output
  aPerson after swapping [1:3] and [4:6]
     Person {id:0, Name:Smith}
     Person {id:4, Name:Babar}
     Person {id:5, Name:ONeil}
     Person {id:3, Name:Xu}
     Person {id:1, Name:Jones}
     Person {id:2, Name:Salinas}
     Person {id:6, Name:Akuneli}
     Person {id:7, Name:Yamamoto}

  */
}
