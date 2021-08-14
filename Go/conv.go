package conv

import (
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/conv/Go"
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
// Int8 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int8ToString parses int8 int8_ to string.
func Int8ToString(int8_ int8) (string_ string) {
	return strconv.FormatInt(int64(int8_), 10)
}

// Int8ToPtr is equivalent to Int8ToPtrInt8(int8_).
func Int8ToPtr(int8_ int8) (ptr_int_ *int8) {
	return Int8ToPtrInt8(int8_)
}

// Int8ToPtrString parses int8 int8_ to *string.
func Int8ToPtrString(int8_ int8) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = strconv.FormatInt(int64(int8_), 10)

	return
}

// Int8ToPtrUint64 parses int8 int8_ to *uint64
func Int8ToPtrUint64(int8_ int8) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int8_)

	return
}

// Int8ToPtrUint32 parses int8 int8_ to *uint32
func Int8ToPtrUint32(int8_ int8) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int8_)

	return
}

// Int8ToPtrUint16 parses int8 int8_ to *uint16
func Int8ToPtrUint16(int8_ int8) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int8_)

	return
}

// Int8ToPtrUint8 parses int8 int8_ to *uint8
func Int8ToPtrUint8(int8_ int8) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int8_)

	return
}

// Int8ToPtrUint parses int8 int8_ to *uint
func Int8ToPtrUint(int8_ int8) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int8_)

	return
}

// Int8ToPtrInt64 parses int8 int8_ to *int64
func Int8ToPtrInt64(int8_ int8) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int8_)

	return
}

// Int8ToPtrInt32 parses int8 int8_ to *int32
func Int8ToPtrInt32(int8_ int8) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int8_)

	return
}

// Int8ToPtrInt16 parses int8 int8_ to *int16
func Int8ToPtrInt16(int8_ int8) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int8_)

	return
}

// Int8ToPtrInt8 parses int8 int8_ to *int8
func Int8ToPtrInt8(int8_ int8) (ptr_int8_ *int8) {
	return &int8_
}

// Int8ToPtrInt parses int8 int8_ to *int
func Int8ToPtrInt(int8_ int8) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int8_)

	return
}

// Int8ToPtrBool parses int8 int8_ to *bool
func Int8ToPtrBool(int8_ int8) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int8_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int8_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int8ToUint64 parses int8 int8_ to uint64.
func Int8ToUint64(int8_ int8) (uint64_ uint64) {
	uint64_ = uint64(int8_)

	return
}

// Int8ToUint32 parses int8 int8_ to uint32.
func Int8ToUint32(int8_ int8) (uint32_ uint32) {
	uint32_ = uint32(int8_)

	return
}

// Int8ToUint16 parses int8 int8_ to uint16.
func Int8ToUint16(int8_ int8) (uint16_ uint16) {
	uint16_ = uint16(int8_)

	return
}

// Int8ToUint8 parses int8 int8_ to uint8.
func Int8ToUint8(int8_ int8) (uint8_ uint8) {
	uint8_ = uint8(int8_)

	return
}

// Int8ToUint parses int8 int8_ to uint.
func Int8ToUint(int8_ int8) (uint_ uint) {
	uint_ = uint(int8_)

	return
}

// Int8ToInt64 parses int8 int8_ to int64.
func Int8ToInt64(int8_ int8) (int64_ int64) {
	int64_ = int64(int8_)

	return
}

// Int8ToInt32 parses int8 int8_ to int32.
func Int8ToInt32(int8_ int8) (int32_ int32) {
	int32_ = int32(int8_)

	return
}

// Int8ToInt16 parses int8 int8_ to int16.
func Int8ToInt16(int8_ int8) (int16_ int16) {
	int16_ = int16(int8_)

	return
}

// Int8ToInt parses int8 int8_ to int.
func Int8ToInt(int8_ int8) (int_ int) {
	int_ = int(int8_)

	return
}

// Int8ToBool parses int8 int8_ to bool.
func Int8ToBool(int8_ int8) (bool_ bool) {
	if int8_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int8_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// IntToString parses int int_ to string.
func IntToString(int_ int) (string_ string) {
	return strconv.FormatInt(int64(int_), 10)
}

// IntToPtr is equivalent to IntToPtrInt(int_).
func IntToPtr(int_ int) (ptr_int_ *int) {
	return IntToPtrInt(int_)
}

// IntToPtrString parses int int_ to *string.
func IntToPtrString(int_ int) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = strconv.FormatInt(int64(int_), 10)

	return
}

// IntToPtrUint64 parses int int_ to *uint64.
func IntToPtrUint64(int_ int) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int_)

	return
}

// IntToPtrUint32 parses int int_ to *uint32.
func IntToPtrUint32(int_ int) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int_)

	return
}

// IntToPtrUint16 parses int int_ to *uint16.
func IntToPtrUint16(int_ int) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int_)

	return
}

// IntToPtrUint8 parses int int_ to *uint8.
func IntToPtrUint8(int_ int) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int_)

	return
}

// IntToPtrUint parses int int_ to *uint.
func IntToPtrUint(int_ int) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int_)

	return
}

// IntToPtrInt64 parses int int_ to *int64.
func IntToPtrInt64(int_ int) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int_)

	return
}

// IntToPtrInt32 parses int int_ to *int32.
func IntToPtrInt32(int_ int) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int_)

	return
}

// IntToPtrInt16 parses int int_ to *int16.
func IntToPtrInt16(int_ int) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int_)

	return
}

// IntToPtrInt8 parses int int_ to *int8.
func IntToPtrInt8(int_ int) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int_)

	return
}

// IntToPtrInt parses int int_ to *int.
func IntToPtrInt(int_ int) (ptr_int_ *int) {
	return &int_
}

// IntToPtrBool parses int int_ to *bool.
func IntToPtrBool(int_ int) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// IntToUint64 parses int int_ to uint64.
func IntToUint64(int_ int) (uint64_ uint64) {
	return uint64(int_)
}

// IntToUint32 parses int int_ to uint32.
func IntToUint32(int_ int) (uint32_ uint32) {
	return uint32(int_)
}

// IntToUint16 parses int int_ to uint16.
func IntToUint16(int_ int) (uint16_ uint16) {
	return uint16(int_)
}

// IntToUint8 parses int int_ to uint8.
func IntToUint8(int_ int) (uint8_ uint8) {
	return uint8(int_)
}

// IntToUint parses int int_ to uint.
func IntToUint(int_ int) (uint_ uint) {
	return uint(int_)
}

// IntToInt64 parses int int_ to int64.
func IntToInt64(int_ int) (int64_ int64) {
	return int64(int_)
}

// IntToInt32 parses int int_ to int32.
func IntToInt32(int_ int) (int32_ int32) {
	return int32(int_)
}

// IntToInt16 parses int int_ to int16.
func IntToInt16(int_ int) (int16_ int16) {
	return int16(int_)
}

// IntToInt8 parses int int_ to int8.
func IntToInt8(int_ int) (int8_ int8) {
	return int8(int_)
}

// IntToBool parses int int_ to bool.
func IntToBool(int_ int) (bool_ bool) {
	if int_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
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
