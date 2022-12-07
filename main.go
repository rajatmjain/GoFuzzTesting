package main

import (
	"errors"
	"fmt"
	"unicode/utf8"

	myFuzz "GoFuzzTesting/gofuzz"

	goFuzz "github.com/google/gofuzz"
)

func main() {
	testString()
}

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

func testString(){
	generatedString := MyStringGenerator('0','ӿ',10)
	fmt.Println("Original string:",generatedString)
	reversed,reversedError := Reverse(generatedString)
	doubleReversed,doubleReversedError := Reverse(reversed)
	if(reversedError==nil && doubleReversedError==nil){
		fmt.Println("Reversed string:",reversed)
		fmt.Println("Double Reversed string:",doubleReversed)
	}else{
		fmt.Println("Reversed error:",reversedError,"Original error:",doubleReversedError)
	}
	fmt.Println("Is the generated string same as the double reversed string?:",doubleReversed==generatedString)
}

func MyStringGenerator(first rune,last rune, maxLength int) string{
	// String Fuzzer
	var A string
	unicodeRange := myFuzz.UnicodeRange{First: first, Last: last, MaxLength: maxLength}
	//unicodeRange := myFuzz.UnicodeRange{First: '0', Last: 'ӿ', MaxLength: 25}
	stringFuzzer := myFuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&A)
    return A
}




func IsPositiveAndEven(i int) bool{
	if(i>0 && i%2==0){
		return true
	}else{
		return false
	}
}


// GoFuzz v/s MyFuzz
func goFuzzGenerator(){
	fmt.Println("Go Fuzzer")
	// String Fuzzer
	var A string
	unicodeRange := goFuzz.UnicodeRange{First: '0', Last: 'ӿ'}
	stringFuzzer := goFuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&A)
    fmt.Println("Fuzzed string: ",A)

	// Integer64 Fuzzer
	var B int64
	int64Fuzzer := goFuzz.New()
	int64Fuzzer.Fuzz(&B)
	fmt.Println("Fuzzed integer64: ",B)

	// Integer32 Fuzzer
	var C int32
	int32Fuzzer := myFuzz.New()
	int32Fuzzer.Fuzz(&C)
	fmt.Println("Fuzzed integer32: ",C)

	//Float64 Fuzzer
	var D float64
	float64Fuzzer := myFuzz.New()
	float64Fuzzer.Fuzz(&D)
	fmt.Println("Fuzzed float64: ",D)

	//Float32 Fuzzer
	var E float32
	float32Fuzzer := myFuzz.New()
	float32Fuzzer.Fuzz(&E)
	fmt.Println("Fuzzed float32: ",E)
}

func myFuzzGenerator(){
	fmt.Println("Schematic Fuzzer")
	// String Fuzzer
	var A string
	unicodeRange := myFuzz.UnicodeRange{First: '0', Last: 'ӿ',MaxLength: 25}
	stringFuzzer := myFuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	stringFuzzer.Fuzz(&A)
    fmt.Println("Fuzzed string: ",A)

	// Integer64 Fuzzer
	var B int64
	int64Schema := myFuzz.Int64Schema{Minimum: 0,Maximum: 10000}
	int64Fuzzer := myFuzz.New().Funcs(int64Schema.CustomInt64FuzzFunc())
	int64Fuzzer.Fuzz(&B)
	fmt.Println("Fuzzed integer64: ",B)

	// Integer32 Fuzzer
	var C int32
	int32Schema := myFuzz.Int32Schema{Minimum: 0,Maximum: 10000}
	int32Fuzzer := myFuzz.New().Funcs(int32Schema.CustomInt32FuzzFunc())
	int32Fuzzer.Fuzz(&C)
	fmt.Println("Fuzzed integer32: ",C)

	//Float64 Fuzzer
	var D float64
	float64Schema := myFuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 9}
	float64Fuzzer := myFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&D)
	fmt.Println("Fuzzed float64: ",D)

	//Float32 Fuzzer
	var E float32
	float32Schema := myFuzz.Float32Schema{Minimum: 0,Maximum: 10000,Precision: 2}
	float32Fuzzer := myFuzz.New().Funcs(float32Schema.CustomFloat32FuzzFunc())
	float32Fuzzer.Fuzz(&E)
	fmt.Println("Fuzzed float32: ",E)
}

