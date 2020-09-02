package p

// T type testing the export
type T struct{ a, b int } // a and b are not exported

// U type testing the export
type U struct {
	A, B int
	c    int
} // A and B are exported only
