package conv

import (
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// strconv
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Atoi is equivalent to strconv.ParseInt(s, 10, 0), converted to type int.
func Atoi(string_ string) (int_ int, err error) {
	return strconv.Atoi(string_)
}

// Itoa is equivalent to strconv.FormatInt(int64(int_), 10).
func Itoa(int_ int) (string_ string) {
	return strconv.Itoa(int_)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Bool conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// BoolToString parses bool bool_ to string.
func BoolToString(bool_ bool) (string_ string) {
	return strconv.FormatBool(bool_)
}

// BoolToPtr is equivalent to BoolToPtrBool(bool_).
func BoolToPtr(bool_ bool) (ptr_bool_ *bool) {
	return BoolToPtrBool(bool_)
}

// BoolToPtrString parses bool bool_ to *string.
func BoolToPtrString(bool_ bool) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = strconv.FormatBool(bool_)

	return
}

// BoolToPtrUint parses bool bool_ to *uint.
func BoolToPtrUint(bool_ bool) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	switch bool_ {
	case false:
		*ptr_uint_ = 0
	case true:
		*ptr_uint_ = 1
	default:
		ptr_uint_ = nil
	}

	return
}

// BoolToPtrInt parses bool bool_ to *int.
func BoolToPtrInt(bool_ bool) (ptr_int_ *int) {
	ptr_int_ = new(int)

	switch bool_ {
	case false:
		*ptr_int_ = 0
	case true:
		*ptr_int_ = 1
	default:
		ptr_int_ = nil
	}

	return
}

// BoolToPtrBool parses bool bool_ to *bool.
func BoolToPtrBool(bool_ bool) (ptr_bool_ *bool) {
	return &bool_
}

// BoolToUint parses bool bool_ to uint.
func BoolToUint(bool_ bool) (uint_ uint) {
	switch bool_ {
	case false:
		uint_ = 0
	case true:
		uint_ = 1
	default:
	}

	return
}

// BoolToInt parses bool bool_ to int.
func BoolToInt(bool_ bool) (int_ int) {
	switch bool_ {
	case false:
		int_ = 0
	case true:
		int_ = 1
	default:
	}

	return
}
