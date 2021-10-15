package conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Uint16 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Uint16ToString parses uint16 uint16_ to string.
func Uint16ToString(uint16_ uint16) (string_ string) {
	string_ = fmt.Sprint(uint16_)

	return
}

// Uint16ToPtr is equivalent to Uint16ToPtrUint16(uint16_).
func Uint16ToPtr(uint16_ uint16) (ptr_uint16_ *uint16) {
	ptr_uint16_ = Uint16ToPtrUint16(uint16_)

	return
}

// Uint16ToPtrString parses uint16 uint16_ to *string.
func Uint16ToPtrString(uint16_ uint16) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(uint16_)

	return
}

// Uint16ToPtrUint64 parses uint16 uint16_ to *uint64.
func Uint16ToPtrUint64(uint16_ uint16) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(uint16_)

	return
}

// Uint16ToPtrUint32 parses uint16 uint16_ to *uint32.
func Uint16ToPtrUint32(uint16_ uint16) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(uint16_)

	return
}

// Uint16ToPtrUint16 parses uint16 uint16_ to *uint16.
func Uint16ToPtrUint16(uint16_ uint16) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(uint16_)

	return
}

// Uint16ToPtrUint8 parses uint16 uint16_ to *uint8.
func Uint16ToPtrUint8(uint16_ uint16) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(uint16_)

	return
}

// Uint16ToPtrUint parses uint16 uint16_ to *uint.
func Uint16ToPtrUint(uint16_ uint16) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(uint16_)

	return
}

// Uint16ToPtrInt64 parses uint16 uint16_ to *int64.
func Uint16ToPtrInt64(uint16_ uint16) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(uint16_)

	return
}

// Uint16ToPtrInt32 parses uint16 uint16_ to *int32.
func Uint16ToPtrInt32(uint16_ uint16) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(uint16_)

	return
}

// Uint16ToPtrInt16 parses uint16 uint16_ to *int16.
func Uint16ToPtrInt16(uint16_ uint16) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(uint16_)

	return
}

// Uint16ToPtrInt8 parses uint16 uint16_ to *int8.
func Uint16ToPtrInt8(uint16_ uint16) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(uint16_)

	return
}

// Uint16ToPtrInt parses uint16 uint16_ to *int.
func Uint16ToPtrInt(uint16_ uint16) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(uint16_)

	return
}

// Uint16ToPtrBool parses uint16 uint16_ to *bool.
func Uint16ToPtrBool(uint16_ uint16) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if uint16_ == 0 {
		*ptr_bool_ = false
	} else if uint16_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Uint16ToUint64 parses uint16 uint16_ to uint64.
func Uint16ToUint64(uint16_ uint16) (uint64_ uint64) {
	uint64_ = uint64(uint16_)

	return
}

// Uint16ToUint32 parses uint16 uint16_ to uint32.
func Uint16ToUint32(uint16_ uint16) (uint32_ uint32) {
	uint32_ = uint32(uint16_)

	return
}

// Uint16ToUint16 returns uint16 uint16_.
func Uint16ToUint16(uint16_ uint16) uint16 {
	return uint16_
}

// Uint16ToUint8 parses uint16 uint16_ to uint8.
func Uint16ToUint8(uint16_ uint16) (uint8_ uint8) {
	uint8_ = uint8(uint16_)

	return
}

// Uint16ToUint parses uint16 uint16_ to uint.
func Uint16ToUint(uint16_ uint16) (uint_ uint) {
	uint_ = uint(uint16_)

	return
}

// Uint16ToInt64 parses uint16 uint16_ to int64.
func Uint16ToInt64(uint16_ uint16) (int64_ int64) {
	int64_ = int64(uint16_)

	return
}

// Uint16ToInt32 parses uint16 uint16_ to int32.
func Uint16ToInt32(uint16_ uint16) (int32_ int32) {
	int32_ = int32(uint16_)

	return
}

// Uint16ToInt16 parses uint16 uint16_ to int16.
func Uint16ToInt16(uint16_ uint16) (int16_ int16) {
	int16_ = int16(uint16_)

	return
}

// Uint16ToInt8 parses uint16 uint16_ to int8.
func Uint16ToInt8(uint16_ uint16) (int8_ int8) {
	int8_ = int8(uint16_)

	return
}

// Uint16ToInt parses uint16 uint16_ to int.
func Uint16ToInt(uint16_ uint16) (int_ int) {
	int_ = int(uint16_)

	return
}

// Uint16ToBool parses uint16 uint16_ to bool.
func Uint16ToBool(uint16_ uint16) (bool_ bool) {
	if uint16_ == 0 {
		bool_ = false
	} else if uint16_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
