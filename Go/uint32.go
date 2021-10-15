package conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Uint32 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Uint32ToString parses uint32 uint32_ to string.
func Uint32ToString(uint32_ uint32) (string_ string) {
	string_ = fmt.Sprint(uint32_)

	return
}

// Uint32ToPtr is equivalent to Uint32ToPtrUint32(uint32_).
func Uint32ToPtr(uint32_ uint32) (ptr_uint32_ *uint32) {
	ptr_uint32_ = Uint32ToPtrUint32(uint32_)

	return
}

// Uint32ToPtrString parses uint32 uint32_ to *string.
func Uint32ToPtrString(uint32_ uint32) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(uint32_)

	return
}

// Uint32ToPtrUint64 parses uint32 uint32_ to *uint64.
func Uint32ToPtrUint64(uint32_ uint32) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(uint32_)

	return
}

// Uint32ToPtrUint32 parses uint32 uint32_ to *uint32.
func Uint32ToPtrUint32(uint32_ uint32) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(uint32_)

	return
}

// Uint32ToPtrUint16 parses uint32 uint32_ to *uint16.
func Uint32ToPtrUint16(uint32_ uint32) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(uint32_)

	return
}

// Uint32ToPtrUint8 parses uint32 uint32_ to *uint8.
func Uint32ToPtrUint8(uint32_ uint32) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(uint32_)

	return
}

// Uint32ToPtrUint parses uint32 uint32_ to *uint.
func Uint32ToPtrUint(uint32_ uint32) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(uint32_)

	return
}

// Uint32ToPtrInt64 parses uint32 uint32_ to *int64.
func Uint32ToPtrInt64(uint32_ uint32) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(uint32_)

	return
}

// Uint32ToPtrInt32 parses uint32 uint32_ to *int32.
func Uint32ToPtrInt32(uint32_ uint32) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(uint32_)

	return
}

// Uint32ToPtrInt16 parses uint32 uint32_ to *int16.
func Uint32ToPtrInt16(uint32_ uint32) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(uint32_)

	return
}

// Uint32ToPtrInt8 parses uint32 uint32_ to *int8.
func Uint32ToPtrInt8(uint32_ uint32) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(uint32_)

	return
}

// Uint32ToPtrInt parses uint32 uint32_ to *int.
func Uint32ToPtrInt(uint32_ uint32) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(uint32_)

	return
}

// Uint32ToPtrBool parses uint32 uint32_ to *bool.
func Uint32ToPtrBool(uint32_ uint32) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if uint32_ == 0 {
		*ptr_bool_ = false
	} else if uint32_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Uint32ToUint64 parses uint32 uint32_ to uint64.
func Uint32ToUint64(uint32_ uint32) (uint64_ uint64) {
	uint64_ = uint64(uint32_)

	return
}

// Uint32ToUint32 returns uint32 uint32_.
func Uint32ToUint32(uint32_ uint32) uint32 {
	return uint32_
}

// Uint32ToUint16 parses uint32 uint32_ to uint16.
func Uint32ToUint16(uint32_ uint32) (uint16_ uint16) {
	uint16_ = uint16(uint32_)

	return
}

// Uint32ToUint8 parses uint32 uint32_ to uint8.
func Uint32ToUint8(uint32_ uint32) (uint8_ uint8) {
	uint8_ = uint8(uint32_)

	return
}

// Uint32ToUint parses uint32 uint32_ to uint.
func Uint32ToUint(uint32_ uint32) (uint_ uint) {
	uint_ = uint(uint32_)

	return
}

// Uint32ToInt64 parses uint32 uint32_ to int64.
func Uint32ToInt64(uint32_ uint32) (int64_ int64) {
	int64_ = int64(uint32_)

	return
}

// Uint32ToInt32 parses uint32 uint32_ to int32.
func Uint32ToInt32(uint32_ uint32) (int32_ int32) {
	int32_ = int32(uint32_)

	return
}

// Uint32ToInt16 parses uint32 uint32_ to int16.
func Uint32ToInt16(uint32_ uint32) (int16_ int16) {
	int16_ = int16(uint32_)

	return
}

// Uint32ToInt8 parses uint32 uint32_ to int8.
func Uint32ToInt8(uint32_ uint32) (int8_ int8) {
	int8_ = int8(uint32_)

	return
}

// Uint32ToInt parses uint32 uint32_ to int.
func Uint32ToInt(uint32_ uint32) (int_ int) {
	int_ = int(uint32_)

	return
}

// Uint32ToBool parses uint32 uint32_ to bool.
func Uint32ToBool(uint32_ uint32) (bool_ bool) {
	if uint32_ == 0 {
		bool_ = false
	} else if uint32_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
