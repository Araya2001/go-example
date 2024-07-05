package main

import (
	"fmt"
	"go-example/example/service"
)

func main() {
	fmt.Println("Using Interface Sampler")

	// Define the simple struct for the example to be used
	simpleStruct := service.SimpleStruct{
		AttrOne:   "attribute one",
		AttrTwo:   "attribute two",
		AttrThree: "attribute three",
	}

	// Define the Interface that's going to be used
	// (It uses pointers because we need to modify the struct from within the implementation)
	simpleStructSampler := service.SimpleStructSampler(&simpleStruct)

	// Modify the simple struct through args, the "_" is an unhandled error, it ignores the value
	_ = simpleStructSampler.ModifySimpleStruct("ab", "", "bc")

	// Get the pointer to the initial struct named "simpleStruct"
	modifiedSimpleStructPtr, _ := simpleStructSampler.GetSimpleStruct()

	// Copy the value of the struct retrieved from the simpleStructSampler.GetSimpleStruct() method call
	modifiedSimpleStruct := *modifiedSimpleStructPtr

	// Compare the outputs visually when running this piece of code (They should be equal,
	// but the modifiedSimpleStruct has a new memory value)
	fmt.Println(modifiedSimpleStruct)
	fmt.Println(simpleStruct)

	// Modify the value-copied struct
	modifiedSimpleStruct.AttrOne = "attribute one modified"

	// Compare again the outputs visually when running this piece of code (They're not equal)
	fmt.Println(modifiedSimpleStruct)
	fmt.Println(simpleStruct)

}
