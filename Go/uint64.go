package Conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import Conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// uint64 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Uint64ToString parses uint64 uint64_ to string.
func Uint64ToString(uint64_ uint64) (string_ string) {
	string_ = fmt.Sprint(uint64_)

	return
}

// Uint64ToPtr is equivalent to Uint64ToPtrUint64(uint64_).
func Uint64ToPtr(uint64_ uint64) (ptr_uint64_ *uint64) {
	ptr_uint64_ = Uint64ToPtrUint64(uint64_)

	return
}

// Uint64ToPtrString parses uint64 uint64_ to *string.
func Uint64ToPtrString(uint64_ uint64) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(uint64_)

	return
}

// Uint64ToPtrUint64 parses uint64 uint64_ to *uint64.
func Uint64ToPtrUint64(uint64_ uint64) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(uint64_)

	return
}

// Uint64ToPtrUint32 parses uint64 uint64_ to *uint32.
func Uint64ToPtrUint32(uint64_ uint64) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(uint64_)

	return
}

// Uint64ToPtrUint16 parses uint64 uint64_ to *uint16.
func Uint64ToPtrUint16(uint64_ uint64) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(uint64_)

	return
}

// Uint64ToPtrUint8 parses uint64 uint64_ to *uint8.
func Uint64ToPtrUint8(uint64_ uint64) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(uint64_)

	return
}

// Uint64ToPtrUint parses uint64 uint64_ to *uint.
func Uint64ToPtrUint(uint64_ uint64) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(uint64_)

	return
}

// Uint64ToPtrInt64 parses uint64 uint64_ to *int64.
func Uint64ToPtrInt64(uint64_ uint64) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(uint64_)

	return
}

// Uint64ToPtrInt32 parses uint64 uint64_ to *int32.
func Uint64ToPtrInt32(uint64_ uint64) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(uint64_)

	return
}

// Uint64ToPtrInt16 parses uint64 uint64_ to *int16.
func Uint64ToPtrInt16(uint64_ uint64) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(uint64_)

	return
}

// Uint64ToPtrInt8 parses uint64 uint64_ to *int8.
func Uint64ToPtrInt8(uint64_ uint64) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(uint64_)

	return
}

// Uint64ToPtrInt parses uint64 uint64_ to *int.
func Uint64ToPtrInt(uint64_ uint64) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(uint64_)

	return
}

// Uint64ToPtrBool parses uint64 uint64_ to *bool.
func Uint64ToPtrBool(uint64_ uint64) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if uint64_ == 0 {
		*ptr_bool_ = false
	} else if uint64_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Uint64ToUint64 returns uint64 uint64_.
func Uint64ToUint64(uint64_ uint64) uint64 {
	return uint64_
}

// Uint64ToUint32 parses uint64 uint64_ to uint32.
func Uint64ToUint32(uint64_ uint64) (uint32_ uint32) {
	uint32_ = uint32(uint64_)

	return
}

// Uint64ToUint16 parses uint64 uint64_ to uint16.
func Uint64ToUint16(uint64_ uint64) (uint16_ uint16) {
	uint16_ = uint16(uint64_)

	return
}

// Uint64ToUint8 parses uint64 uint64_ to uint8.
func Uint64ToUint8(uint64_ uint64) (uint8_ uint8) {
	uint8_ = uint8(uint64_)

	return
}

// Uint64ToUint parses uint64 uint64_ to uint.
func Uint64ToUint(uint64_ uint64) (uint_ uint) {
	uint_ = uint(uint64_)

	return
}

// Uint64ToInt64 parses uint64 uint64_ to int64.
func Uint64ToInt64(uint64_ uint64) (int64_ int64) {
	int64_ = int64(uint64_)

	return
}

// Uint64ToInt32 parses uint64 uint64_ to int32.
func Uint64ToInt32(uint64_ uint64) (int32_ int32) {
	int32_ = int32(uint64_)

	return
}

// Uint64ToInt16 parses uint64 uint64_ to int16.
func Uint64ToInt16(uint64_ uint64) (int16_ int16) {
	int16_ = int16(uint64_)

	return
}

// Uint64ToInt8 parses uint64 uint64_ to int8.
func Uint64ToInt8(uint64_ uint64) (int8_ int8) {
	int8_ = int8(uint64_)

	return
}

// Uint64ToInt parses uint64 uint64_ to int.
func Uint64ToInt(uint64_ uint64) (int_ int) {
	int_ = int(uint64_)

	return
}

// Uint64ToBool parses uint64 uint64_ to bool.
func Uint64ToBool(uint64_ uint64) (bool_ bool) {
	if uint64_ == 0 {
		bool_ = false
	} else if uint64_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
