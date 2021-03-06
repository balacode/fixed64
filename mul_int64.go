// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 18:46:07 25BE78                 priveda/fixed64/[mul_int64.go]
// -----------------------------------------------------------------------------

package fixed64

// MulInt64 multiplies a fixed-point number by one or more 64-bit integers
// and returns the result. The original number is not changed.
func (n Fixed64) MulInt64(nums ...int64) Fixed64 {
	for _, num := range nums {
		n = n.Mul(Fixed64{num * 1E4})
	}
	return n
}

//end
