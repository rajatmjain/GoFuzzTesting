package main

import (
	"testing"

	schemaBasedInputGenerator "github.com/rajatjain007/GoFuzzTesting/helper"
)

// func FuzzIsPositiveAndEven(f *testing.F) {
//     f.Fuzz(func(t *testing.T, input int){
//         isPositive,isEven := IsPositiveAndEven(input)
//         if(isPositive!=true || isEven!=true){
//             t.Errorf("Error!")
//         }

//     })
// }

func TestIsPositiveAndEven(t *testing.T) {
    value := schemaBasedInputGenerator.ExampleCustomInt(schemaBasedInputGenerator.CustomInt{1,10000})
    isPositive,isEven := IsPositiveAndEven(value)
        if (isPositive != true && isEven != true) {
                t.Errorf("Error!")
        }
}
