package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("/// Output: exercisePointerChangeColor() ///")
	exercisePointerChangeColor()

	fmt.Println("\n/// Output: exercisePointerToPointer() ///")
	exercisePointerToPointer()

}

func exercisePointerChangeColor() {
	var color string = ""
	var colorPtr = &color // assign address of color to colorPtr.
	// colorPtr implicitly has type, *string.
	// ie colorPtr is a string pointer that
	//  refers to address of color.

	// pass by address
	changeColor(&color)
	fmt.Printf("after first call : color =%s\n", color)

	//
	changeColor(colorPtr)
	fmt.Printf("after second call : color =%s\n", color)
}

var aryColor = [...]string{"red", "green", "blue", "purple", "indigo", "yellow", "black", "gray", "white"}

/**
 *  This function exemplifies pass by reference
 *  and mutates the :param:colorPtr
 *
 *  @param colorPtr :*string [in/out parameter]
 */
func changeColor(colorPtr *string) {
	var newColorIndex = rand.Int() % len(aryColor)
	*colorPtr = aryColor[newColorIndex]
}

func exercisePointerToPointer() {
	var numVarAtLoc00 = 100
	var ptrToLoc00 = &numVarAtLoc00
	var prtToPtrToLoc00 = &ptrToLoc00

	fmt.Printf(`
Initial State:
numVarAtLoc00   =%16d    // value in   loc00
&numVarAtLoc00  =%16d    // address of loc00
ptrToLoc00      =%16d    // address of loc00
prtToPtrToLoc00 =%16d    // address of ptrToLoc00
*ptrToLoc00     =%16d    // value in   loc00  // example of following ptr to value
**ptrToPtrToLoc =%16d    // value in   loc00  // example of following ptr to ptr to value
`, numVarAtLoc00, &numVarAtLoc00, ptrToLoc00, prtToPtrToLoc00, *ptrToLoc00, **prtToPtrToLoc00)
}
