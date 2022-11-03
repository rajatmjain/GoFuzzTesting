package schemaBasedInputGenerator

import (
	"math"
	"math/rand"
	"time"
)

// CUSTOM DATA STRUCTURE SCHEMAS
type CustomInt struct{
    Min int
    Max int
}

type CustomFloat64 struct {
	Min float64
	Max float64
	Precision uint
}

type CustomFloat32 struct {
	Min float32
	Max float32
	Precision uint
}

type CustomString struct{
	First rune
	Last rune
	MaxLength int
  
}

// Custom Integer Generator
func ExampleCustomInt(schema CustomInt) int{
	rand.Seed(time.Now().UnixNano())
    min := schema.Min
    max := schema.Max
    result := rand.Intn(max-min) + min
    return result
}

// Custom Float64 Generator
func ExampleCustomFloat64(schema CustomFloat64) float64{
	rand.Seed(time.Now().UnixNano())
	min := schema.Min
    max := schema.Max
	precision := schema.Precision
	ratio := math.Pow(10, float64(precision))
	result := min + rand.Float64() * (max - min)
	return math.Round(result*ratio)/ratio
}

// Custom Float Generator
func ExampleCustomFloat32(schema CustomFloat32) float32{
	rand.Seed(time.Now().UnixNano())
	min := schema.Min
    max := schema.Max
	precision := schema.Precision
	ratio := math.Pow(10, float64(precision))
	result := min + rand.Float32() * (max - min)
	return float32(math.Round(float64(result)*ratio))/float32(ratio)
}

// Custom String Generator
// func ExampleCustomString(schema CustomString) {
// 	first := CustomString.First
// 	last := CustomString.Last
// 	maxLength := CustomString.MaxLength
	
// }