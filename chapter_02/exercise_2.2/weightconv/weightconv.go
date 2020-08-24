// Package weightconv performs Kilogram and Pound conversions.
package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }

// KToP converts a Kilogram weight to Pound.
func KToP(k Kilogram) Pound { return Pound(k * 2.20462) }

// PToK converts a Pound weight to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }
