package conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Uint8 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Uint8ToString parses uint8 uint8_ to string.
func Uint8ToString(uint8_ uint8) (string_ string) {
	string_ = fmt.Sprint(uint8_)

	return
}

// Uint8ToPtr is equivalent to Uint8ToPtrUint8(uint8_).
func Uint8ToPtr(uint8_ uint8) (ptr_uint8_ *uint8) {
	ptr_uint8_ = Uint8ToPtrUint8(uint8_)

	return
}

// Uint8ToPtrString parses uint8 uint8_ to *string.
func Uint8ToPtrString(uint8_ uint8) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(uint8_)

	return
}

// Uint8ToPtrUint64 parses uint8 uint8_ to *uint64.
func Uint8ToPtrUint64(uint8_ uint8) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(uint8_)

	return
}

// Uint8ToPtrUint32 parses uint8 uint8_ to *uint32.
func Uint8ToPtrUint32(uint8_ uint8) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(uint8_)

	return
}

// Uint8ToPtrUint16 parses uint8 uint8_ to *uint16.
func Uint8ToPtrUint16(uint8_ uint8) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(uint8_)

	return
}

// Uint8ToPtrUint8 parses uint8 uint8_ to *uint8.
func Uint8ToPtrUint8(uint8_ uint8) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8_

	return
}

// Uint8ToPtrUint parses uint8 uint8_ to *uint.
func Uint8ToPtrUint(uint8_ uint8) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(uint8_)

	return
}

// Uint8ToPtrInt64 parses uint8 uint8_ to *int64.
func Uint8ToPtrInt64(uint8_ uint8) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(uint8_)

	return
}

// Uint8ToPtrInt32 parses uint8 uint8_ to *int32.
func Uint8ToPtrInt32(uint8_ uint8) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(uint8_)

	return
}

// Uint8ToPtrInt16 parses uint8 uint8_ to *int16.
func Uint8ToPtrInt16(uint8_ uint8) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(uint8_)

	return
}

// Uint8ToPtrInt8 parses uint8 uint8_ to *int8.
func Uint8ToPtrInt8(uint8_ uint8) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(uint8_)

	return
}

// Uint8ToPtrInt parses uint8 uint8_ to *int.
func Uint8ToPtrInt(uint8_ uint8) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(uint8_)

	return
}

// Uint8ToPtrBool parses uint8 uint8_ to *bool.
func Uint8ToPtrBool(uint8_ uint8) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if uint8_ == 0 {
		*ptr_bool_ = false
	} else if uint8_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Uint8ToUint64 parses uint8 uint8_ to uint64.
func Uint8ToUint64(uint8_ uint8) (uint64_ uint64) {
	uint64_ = uint64(uint8_)

	return
}

// Uint8ToUint32 parses uint8 uint8_ to uint32.
func Uint8ToUint32(uint8_ uint8) (uint32_ uint32) {
	uint32_ = uint32(uint8_)

	return
}

// Uint8ToUint16 parses uint8 uint8_ to uint16.
func Uint8ToUint16(uint8_ uint8) (uint16_ uint16) {
	uint16_ = uint16(uint8_)

	return
}

// Uint8ToUint8 returns uint8 uint8_.
func Uint8ToUint8(uint8_ uint8) uint8 {
	return uint8_
}

// Uint8ToUint parses uint8 uint8_ to uint.
func Uint8ToUint(uint8_ uint8) (uint_ uint) {
	uint_ = uint(uint8_)

	return
}

// Uint8ToInt64 parses uint8 uint8_ to int64.
func Uint8ToInt64(uint8_ uint8) (int64_ int64) {
	int64_ = int64(uint8_)

	return
}

// Uint8ToInt32 parses uint8 uint8_ to int32.
func Uint8ToInt32(uint8_ uint8) (int32_ int32) {
	int32_ = int32(uint8_)

	return
}

// Uint8ToInt16 parses uint8 uint8_ to int16.
func Uint8ToInt16(uint8_ uint8) (int16_ int16) {
	int16_ = int16(uint8_)

	return
}

// Uint8ToInt8 parses uint8 uint8_ to int8.
func Uint8ToInt8(uint8_ uint8) (int8_ int8) {
	int8_ = int8(uint8_)

	return
}

// Uint8ToInt parses uint8 uint8_ to int.
func Uint8ToInt(uint8_ uint8) (int_ int) {
	int_ = int(uint8_)

	return
}

// Uint8ToBool parses uint8 uint8_ to bool.
func Uint8ToBool(uint8_ uint8) (bool_ bool) {
	if uint8_ == 0 {
		bool_ = false
	} else if uint8_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
