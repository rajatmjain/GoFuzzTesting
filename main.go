package main

import (
	"fmt"

	fuzz "github.com/google/gofuzz"
)

func main() {
	// String Fuzzer
	var A string
	unicodeRange := fuzz.UnicodeRange{First: 'A', Last: 'z',MaxLength: 25}
	stringFuzzer := fuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&A)
    fmt.Println("Random string: ",A)

	// Integer64 Fuzzer
	var B int64
	int64Schema := fuzz.Int64Schema{Minimum: 0,Maximum: 10000}
	int64Fuzzer := fuzz.New().Funcs(int64Schema.CustomInt64FuzzFunc())
	int64Fuzzer.Fuzz(&B)
	fmt.Println("Random integer64: ",B)

	// Integer32 Fuzzer
	var C int32
	int32Schema := fuzz.Int32Schema{Minimum: 0,Maximum: 10000}
	int32Fuzzer := fuzz.New().Funcs(int32Schema.CustomInt32FuzzFunc())
	int32Fuzzer.Fuzz(&C)
	fmt.Println("Random integer32: ",C)

	// Float64 Fuzzer
	var D float64
	float64Schema := fuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 0}
	float64Fuzzer := fuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&D)
	fmt.Println("Random float64: ",D)

	// Float32 Fuzzer
	var E float32
	float32Schema := fuzz.Float32Schema{Minimum: 0,Maximum: 10000,Precision: 0}
	float32Fuzzer := fuzz.New().Funcs(float32Schema.CustomFloat32FuzzFunc())
	float32Fuzzer.Fuzz(&E)
	fmt.Println("Random float32: ",E)
	

} 



