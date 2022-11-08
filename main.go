package main

import (
	"fmt"

	fuzz "github.com/google/gofuzz"
)

func main() {
    var A string
	var B int64
	unicodeRange := fuzz.UnicodeRange{First: 'A', Last: 'z',MaxLength: 25}
	f := fuzz.New().Funcs(unicodeRange.CustomStringFuzzFunc())
	f.Fuzz(&A) 
    fmt.Println(A)
	int64Schema := fuzz.Int64Schema{Minimum: 0,Maximum: 10000}
	f1 := fuzz.New().Funcs(int64Schema.CustomInt64FuzzFunc())
	f1.Fuzz(&B)
	fmt.Println(B)
} 



