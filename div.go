// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 00DB4E                       priveda/fixed64/[div.go]
// -----------------------------------------------------------------------------

package fixed64

// Div divides a fixed-point number by one or more other fixed-point
// numbers and returns the result. The original number is not changed.
func (n Fixed64) Div(nums ...Fixed64) Fixed64 {
	for _, num := range nums {
		n.i64 *= 1E4
		n.i64 /= num.i64
	}
	return n
}

//end
