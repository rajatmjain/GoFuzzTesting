package main

import (
	"testing"
	"unicode/utf8"
)

/////////////////////////////////////////////////////////////
// Unit Test
/////////////////////////////////////////////////////////////
// func TestReverse(t *testing.T) {
//     testcases := []struct {
//         in, want string
//     }{
//         {"Hello, world", "dlrow ,olleH"},
//         {" ", " "},
//         {"!12345", "54321!"},
//     }
//     for _, tc := range testcases {
//         rev := Reverse(tc.in)
//         if rev != tc.want {
//                 t.Errorf("Reverse: %q, want %q", rev, tc.want)
//         }
//     }
// }

/////////////////////////////////////////////////////////////
// Naive FuzzReverse
/////////////////////////////////////////////////////////////
// func FuzzReverse(f *testing.F) {
//     testcases := []string{"Hello, world", " ", "!12345"}
//     for _, tc := range testcases {
//         f.Add(tc)  // Use f.Add to provide a seed corpus
//     }
//     f.Fuzz(func(t *testing.T, orig string) {
//         rev := Reverse(orig)
//         doubleRev := Reverse(rev)
//         t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
//         if orig != doubleRev {
//             t.Errorf("Before: %q, after: %q", orig, doubleRev)
//         }
//         if utf8.ValidString(orig) && !utf8.ValidString(rev) {
//             t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
//         }
//     })
// }

/////////////////////////////////////////////////////////////
// Optimal FuzzReverse
/////////////////////////////////////////////////////////////
func FuzzReverse(f *testing.F) {
    testcases := []string {"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}

/////////////////////////////////////////////////////////////
// Largest number fuzz test
/////////////////////////////////////////////////////////////
// func FuzzLargestNumber(f *testing.F) {
//     testcases := [][]int {[]int{99,1,52,6,101,4},[]int{9,101,4,2,6,4,4}}
//     for _, tc := range testcases {
//         f.Add(tc)  // Use f.Add to provide a seed corpus
//     }

//     f.Fuzz(func(t *testing.T, input []int){
//         largest := LargestNumber(input)
//         actual := 101
//         if(largest!=actual){
//             t.Errorf("Received: %d,Actual: %d",largest,actual)
//         }

//     })
// }

/////////////////////////////////////////////////////////////
// Compare 100 dummy fuzz test
/////////////////////////////////////////////////////////////
func FuzzCompare100(f *testing.F) {
    testcases := []int{100,2,349,100,492,2}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }

    f.Fuzz(func(t *testing.T, input int){
        res := Compare100(input)
        if(input!=100 && res==true){
            t.Errorf("wrong")
        }
    })
}