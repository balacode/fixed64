// -----------------------------------------------------------------------------
// (c) admin@priveda.com                                            License: MIT
// :v: 2019-05-14 19:01:20 6A9D2A            priveda/fixed64/[unmarshal_json.go]
// -----------------------------------------------------------------------------

package fixed64

import (
	"encoding/json"
)

// UnmarshalJSON unmarshals a JSON description of Fixed64.
// This method alters the number's value.
func (n *Fixed64) UnmarshalJSON(data []byte) error {
	//  ^  don't remove pointer receiver, it is necessary
	if n == nil {
		return mod.Error(ENilReceiver)
	}
	var (
		num float64
		err = json.Unmarshal(data, &num)
	)
	if err != nil {
		return mod.Error(err)
	}
	n.i64 = int64(num * 1E4)
	return nil
}

//end
