// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 352FC8                       priveda/fixed64/[add.go]
// -----------------------------------------------------------------------------

package fixed64

// Add adds one or more fixed-point numbers and returns a new Fixed64.
// The original number on which Add() is called is not changed.
// If there is an overflow, sets the number's internal value to
// math.MaxInt64 (or math.minInt64+1 when the overflow is negative).
func (n Fixed64) Add(nums ...Fixed64) Fixed64 {
	for _, num := range nums {
		var (
			a = n.i64
			b = num.i64
			c = a + b
		)
		// check for overflow
		if c < minInt64 || (a < 0 && b < 0 && b < (minInt64-a)) {
			return fixed64Overflow(true, EOverflow, ": ", a, " + ", b)
		}
		if c > maxInt64 || (a > 0 && b > 0 && b > (maxInt64-a)) {
			return fixed64Overflow(false, EOverflow, ": ", a, " + ", b)
		}
		n.i64 = c
	}
	return n
}

//end
