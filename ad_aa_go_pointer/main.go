package main

import (
	"fmt"
	"math/rand"
)

func main() {
	exercisePointer()
}

func exercisePointer() {
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
