package Conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import Conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Bool conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// BoolToString parses bool bool_ to string.
func BoolToString(bool_ bool) (string_ string) {
	string_ = fmt.Sprint(bool_)

	return
}

// BoolToPtr is equivalent to BoolToPtrBool(bool_).
func BoolToPtr(bool_ bool) (ptr_bool_ *bool) {
	ptr_bool_ = BoolToPtrBool(bool_)

	return
}

// BoolToPtrString parses bool bool_ to *string.
func BoolToPtrString(bool_ bool) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(bool_)

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
	ptr_bool_ = new(bool)

	*ptr_bool_ = bool(bool_)

	return
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

// BoolToBool returns bool bool_.
func BoolToBool(bool_ bool) bool {
	return bool_
}
