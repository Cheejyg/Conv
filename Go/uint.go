package Conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import Conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Uint conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// UintToString parses uint uint_ to string.
func UintToString(uint_ uint) (string_ string) {
	string_ = fmt.Sprint(uint_)

	return
}

// UintToPtr is equivalent to UintToPtrUint(uint_).
func UintToPtr(uint_ uint) (ptr_uint_ *uint) {
	ptr_uint_ = UintToPtrUint(uint_)
	return
}

// UintToPtrString parses uint uint_ to *string.
func UintToPtrString(uint_ uint) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(uint_)

	return
}

// UintToPtrUint64 parses uint uint_ to *uint64.
func UintToPtrUint64(uint_ uint) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(uint_)

	return
}

// UintToPtrUint32 parses uint uint_ to *uint32.
func UintToPtrUint32(uint_ uint) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(uint_)

	return
}

// UintToPtrUint16 parses uint uint_ to *uint16.
func UintToPtrUint16(uint_ uint) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(uint_)

	return
}

// UintToPtrUint8 parses uint uint_ to *uint8.
func UintToPtrUint8(uint_ uint) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(uint_)

	return
}

// UintToPtrUint parses uint uint_ to *uint.
func UintToPtrUint(uint_ uint) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(uint_)

	return
}

// UintToPtrInt64 parses uint uint_ to *int64.
func UintToPtrInt64(uint_ uint) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(uint_)

	return
}

// UintToPtrInt32 parses uint uint_ to *int32.
func UintToPtrInt32(uint_ uint) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(uint_)

	return
}

// UintToPtrInt16 parses uint uint_ to *int16.
func UintToPtrInt16(uint_ uint) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(uint_)

	return
}

// UintToPtrInt8 parses uint uint_ to *int8.
func UintToPtrInt8(uint_ uint) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(uint_)

	return
}

// UintToPtrInt parses uint uint_ to *int.
func UintToPtrInt(uint_ uint) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(uint_)

	return
}

// UintToPtrBool parses uint uint_ to *bool.
func UintToPtrBool(uint_ uint) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if uint_ == 0 {
		*ptr_bool_ = false
	} else if uint_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// UintToUint64 parses uint uint_ to uint64.
func UintToUint64(uint_ uint) (uint64_ uint64) {
	uint64_ = uint64(uint_)

	return
}

// UintToUint32 parses uint uint_ to uint32.
func UintToUint32(uint_ uint) (uint32_ uint32) {
	uint32_ = uint32(uint_)

	return
}

// UintToUint16 parses uint uint_ to uint16.
func UintToUint16(uint_ uint) (uint16_ uint16) {
	uint16_ = uint16(uint_)

	return
}

// UintToUint8 parses uint uint_ to uint8.
func UintToUint8(uint_ uint) (uint8_ uint8) {
	uint8_ = uint8(uint_)

	return
}

// UintToUint returns uint uint_.
func UintToUint(uint_ uint) uint {
	return uint_
}

// UintToInt64 parses uint uint_ to int64.
func UintToInt64(uint_ uint) (int64_ int64) {
	int64_ = int64(uint_)

	return
}

// UintToInt32 parses uint uint_ to int32.
func UintToInt32(uint_ uint) (int32_ int32) {
	int32_ = int32(uint_)

	return
}

// UintToInt16 parses uint uint_ to int16.
func UintToInt16(uint_ uint) (int16_ int16) {
	int16_ = int16(uint_)

	return
}

// UintToInt8 parses uint uint_ to int8.
func UintToInt8(uint_ uint) (int8_ int8) {
	int8_ = int8(uint_)

	return
}

// UintToInt parses uint uint_ to int.
func UintToInt(uint_ uint) (int_ int) {
	int_ = int(uint_)

	return
}

// UintToBool parses uint uint_ to bool.
func UintToBool(uint_ uint) (bool_ bool) {
	if uint_ == 0 {
		bool_ = false
	} else if uint_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
