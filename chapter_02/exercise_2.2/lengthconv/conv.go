package lengthconv

// FToM converts a Feet length to Meter.
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// MToF converts a Meter length to Feet.
func MToF(m Meter) Feet { return Feet(m * 3.8) }
