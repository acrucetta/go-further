package data

import (
	"fmt"
	"strconv"
)

// Declare a custom Runtime type, which has the underlying type int32 (the same as our
// Movie struct field).
type Runtime int32

// Implement a MarshalJSON() method on the Runtime type so that it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the movie
// runtime (in our case, it will return a string in the format "<runtime> mins").
func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quotedJSONValue := strconv.Quote(jsonValue)

	// Convert the quoted string value to a byte slice and return it.
	return []byte(quotedJSONValue), nil
}

// Implement an UnmarshalJSON() method on the Runtime type so that it satisfies the
// json.Unmarshaler interface. This will allow us to accept strings like "107 mins" or just numbers.
func (r *Runtime) UnmarshalJSON(data []byte) error {
	// Remove quotes if present
	s, err := strconv.Unquote(string(data))
	if err != nil {
		// If unquoting fails, maybe it's a raw number, try to parse as int
		var i int32
		if _, err := fmt.Sscanf(string(data), "%d", &i); err == nil {
			*r = Runtime(i)
			return nil
		}
		return fmt.Errorf("invalid runtime format: %s", string(data))
	}
	// Try to parse the string in the format "<number> mins"
	var i int32
	n, err := fmt.Sscanf(s, "%d mins", &i)
	if err == nil && n == 1 {
		*r = Runtime(i)
		return nil
	}
	// Try to parse as a plain number string
	n, err = fmt.Sscanf(s, "%d", &i)
	if err == nil && n == 1 {
		*r = Runtime(i)
		return nil
	}
	return fmt.Errorf("invalid runtime format: %s", s)
}
