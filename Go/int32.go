package conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int32 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int32ToString parses int32 int32_ to string.
func Int32ToString(int32_ int32) (string_ string) {
	string_ = fmt.Sprint(int32_)

	return
}

// Int32ToPtr is equivalent to Int32ToPtrInt32(int32_).
func Int32ToPtr(int32_ int32) (ptr_int32_ *int32) {
	ptr_int32_ = Int32ToPtrInt32(int32_)

	return
}

// Int32ToPtrString parses int32 int32_ to *string.
func Int32ToPtrString(int32_ int32) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int32_)

	return
}

// Int32ToPtrUint64 parses int32 int32_ to *uint64.
func Int32ToPtrUint64(int32_ int32) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int32_)

	return
}

// Int32ToPtrUint32 parses int32 int32_ to *uint32.
func Int32ToPtrUint32(int32_ int32) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int32_)

	return
}

// Int32ToPtrUint16 parses int32 int32_ to *uint16.
func Int32ToPtrUint16(int32_ int32) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int32_)

	return
}

// Int32ToPtrUint8 parses int32 int32_ to *uint8.
func Int32ToPtrUint8(int32_ int32) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int32_)

	return
}

// Int32ToPtrUint parses int32 int32_ to *uint.
func Int32ToPtrUint(int32_ int32) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int32_)

	return
}

// Int32ToPtrInt64 parses int32 int32_ to *int64.
func Int32ToPtrInt64(int32_ int32) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int32_)

	return
}

// Int32ToPtrInt32 parses int32 int32_ to *int32.
func Int32ToPtrInt32(int32_ int32) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int32_)

	return
}

// Int32ToPtrInt16 parses int32 int32_ to *int16.
func Int32ToPtrInt16(int32_ int32) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int32_)

	return
}

// Int32ToPtrInt8 parses int32 int32_ to *int8.
func Int32ToPtrInt8(int32_ int32) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int32_)

	return
}

// Int32ToPtrInt parses int32 int32_ to *int.
func Int32ToPtrInt(int32_ int32) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int32_)

	return
}

// Int32ToPtrBool parses int32 int32_ to *bool.
func Int32ToPtrBool(int32_ int32) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int32_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int32_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int32ToUint64 parses int32 int32_ to uint64.
func Int32ToUint64(int32_ int32) (uint64_ uint64) {
	uint64_ = uint64(int32_)

	return
}

// Int32ToUint32 parses int32 int32_ to uint32.
func Int32ToUint32(int32_ int32) (uint32_ uint32) {
	uint32_ = uint32(int32_)

	return
}

// Int32ToUint16 parses int32 int32_ to uint16.
func Int32ToUint16(int32_ int32) (uint16_ uint16) {
	uint16_ = uint16(int32_)

	return
}

// Int32ToUint8 parses int32 int32_ to uint8.
func Int32ToUint8(int32_ int32) (uint8_ uint8) {
	uint8_ = uint8(int32_)

	return
}

// Int32ToUint parses int32 int32_ to uint.
func Int32ToUint(int32_ int32) (uint_ uint) {
	uint_ = uint(int32_)

	return
}

// Int32ToInt64 parses int32 int32_ to int64.
func Int32ToInt64(int32_ int32) (int64_ int64) {
	int64_ = int64(int32_)

	return
}

// Int32ToInt32 returns int32 int32_.
func Int32ToInt32(int32_ int32) int32 {
	return int32_
}

// Int32ToInt16 parses int32 int32_ to int16.
func Int32ToInt16(int32_ int32) (int16_ int16) {
	int16_ = int16(int32_)

	return
}

// Int32ToInt8 parses int32 int32_ to int8.
func Int32ToInt8(int32_ int32) (int8_ int8) {
	int8_ = int8(int32_)

	return
}

// Int32ToInt parses int32 int32_ to int.
func Int32ToInt(int32_ int32) (int_ int) {
	int_ = int(int32_)

	return
}

// Int32ToBool parses int32 int32_ to bool.
func Int32ToBool(int32_ int32) (bool_ bool) {
	if int32_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int32_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
