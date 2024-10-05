package satisfactory

import "time"

// CustomTime is a custom type for handling the specific time format in the JSON response.
type CustomTime struct {
	time.Time
}

// UnmarshalJSON parses the custom time format.
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Remove the surrounding quotes
	str := string(b[1 : len(b)-1])
	// Parse the time using the custom format
	t, err := time.Parse("2006.01.02-15.04.05", str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}
