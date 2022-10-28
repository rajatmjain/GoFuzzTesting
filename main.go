package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"unicode/utf8"

	fuzz "github.com/google/gofuzz"
)

func main() {
    // input := "The quick brown fox jumped over the lazy dog"
    // rev, revErr := Reverse(input)
    // doubleRev, doubleRevErr := Reverse(rev)
    // fmt.Printf("original: %q\n", input)
    // fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    // fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
    //ExampleSimple()
    //ExampleDate()
    //ExampleCustomInt(CustomInt{begin: 1,end: 10})
    //ExampleCustomString()
	fmt.Println(ExampleCustomFloat64(CustomFloat64{min: -20.459,max: 100.67,precision: 6}))
    
/////////////////////////////////////////////////////////////
// Largest number
/////////////////////////////////////////////////////////////
    // a := []int{34,12,91,20,0} 
    // b := LargestNumber(a)
    // fmt.Println(b)

/////////////////////////////////////////////////////////////
// Compare 100 
/////////////////////////////////////////////////////////////
    // a:= Compare100(100)
    // fmt.Println(a)
}

// Byte approach
// func Reverse(s string) string {
//     b := []byte(s)
//     for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
//         b[i], b[j] = b[j], b[i]
//     }
//     return string(b)
// }

// Rune approach
func Reverse(s string) (string,error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r),nil
}

func LargestNumber(input []int) int{
    sort.Sort(sort.Reverse(sort.IntSlice(input)))
    return input[0]
}

func Compare100(input int) bool{
    return input==100
}

func ExampleSimple() {
	type MyType struct {
		A string
		B string
		C int
		D struct {
			E float64
		}
	}

	f := fuzz.New()
	object := MyType{}

	uniqueObjects := map[MyType]int{}

	for i := 0; i < 10; i++ {
		f.Fuzz(&object)
		uniqueObjects[object]++
	}
    fmt.Println(uniqueObjects)
	fmt.Printf("Got %v unique objects.\n", len(uniqueObjects))
	// Output:
	// Got 10 unique objects.
}

func ExampleDate(){
    f := fuzz.New()
    var day uint
    var month uint
    var year uint 

    f.Fuzz(&day)
    f.Fuzz(&month)
    f.Fuzz(&year)

    fmt.Println(day)
    fmt.Println(month)
    fmt.Println(year)
}


// Custom primitive datatypes
type CustomString struct{
    
}

type CustomInt struct{
    min int
    max int
}

type CustomFloat64 struct {
	min float64
	max float64
	precision uint
}


// Custom Integer
func ExampleCustomInt(schema CustomInt) int{
    min := schema.min
    max := schema.max
    result := rand.Intn(max-min) + min
    return result
}

func ExampleCustomFloat64(schema CustomFloat64) float64{
	min := schema.min
    max := schema.max
	precision := schema.precision
	ratio := math.Pow(10, float64(precision))
	result := min + rand.Float64() * (max - min)
	return math.Round(result*ratio)/ratio
}



// Custom string, custom float