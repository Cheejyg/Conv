package Conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import Conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// IntToString parses int int_ to string.
func IntToString(int_ int) (string_ string) {
	string_ = fmt.Sprint(int_)
	return
}

// IntToPtr is equivalent to IntToPtrInt(int_).
func IntToPtr(int_ int) (ptr_int_ *int) {
	ptr_int_ = IntToPtrInt(int_)

	return
}

// IntToPtrString parses int int_ to *string.
func IntToPtrString(int_ int) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int_)

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
	ptr_int_ = new(int)

	*ptr_int_ = int(int_)

	return
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
	uint64_ = uint64(int_)

	return
}

// IntToUint32 parses int int_ to uint32.
func IntToUint32(int_ int) (uint32_ uint32) {
	uint32_ = uint32(int_)

	return
}

// IntToUint16 parses int int_ to uint16.
func IntToUint16(int_ int) (uint16_ uint16) {
	uint16_ = uint16(int_)

	return
}

// IntToUint8 parses int int_ to uint8.
func IntToUint8(int_ int) (uint8_ uint8) {
	uint8_ = uint8(int_)

	return
}

// IntToUint parses int int_ to uint.
func IntToUint(int_ int) (uint_ uint) {
	uint_ = uint(int_)

	return
}

// IntToInt64 parses int int_ to int64.
func IntToInt64(int_ int) (int64_ int64) {
	int64_ = int64(int_)

	return
}

// IntToInt32 parses int int_ to int32.
func IntToInt32(int_ int) (int32_ int32) {
	int32_ = int32(int_)

	return
}

// IntToInt16 parses int int_ to int16.
func IntToInt16(int_ int) (int16_ int16) {
	int16_ = int16(int_)

	return
}

// IntToInt8 parses int int_ to int8.
func IntToInt8(int_ int) (int8_ int8) {
	int8_ = int8(int_)

	return
}

// IntToInt returns int int_.
func IntToInt(int_ int) int {
	return int_
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
