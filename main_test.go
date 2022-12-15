package main

// func FuzzReverse(f *testing.F) {
//     testcases := []string{"Hello World!","This is fuzz testing!!"}

//     for _, tc := range testcases {
//         f.Add(tc) // Use f.Add to provide a seed corpus
//     }
//     f.Fuzz(func(t *testing.T, orig string) {
//         rev, err1 := Reverse(orig)
//         if err1 != nil {
//             return
//         }
//         doubleRev, err2 := Reverse(rev)
//         if err2 != nil {
//             return
//         }
//         if orig != doubleRev {
//             t.Errorf("Before: %q, after: %q", orig, doubleRev)
//         }
//         if utf8.ValidString(orig) && !utf8.ValidString(rev) {
//             t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
//         }
//     })
// }

// func FuzzIsPositiveAndEven(f *testing.F) {
//     f.Fuzz(func(t *testing.T, i int){
//         ans := IsPositiveAndEven(i)
//         if(!ans){
//             t.Errorf("No")
//         }
//     })
// }