package Conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import Conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int8 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int8ToString parses int8 int8_ to string.
func Int8ToString(int8_ int8) (string_ string) {
	string_ = fmt.Sprint(int8_)

	return
}

// Int8ToPtr is equivalent to Int8ToPtrInt8(int8_).
func Int8ToPtr(int8_ int8) (ptr_int_ *int8) {
	ptr_int_ = Int8ToPtrInt8(int8_)

	return
}

// Int8ToPtrString parses int8 int8_ to *string.
func Int8ToPtrString(int8_ int8) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int8_)

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
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int8_)

	return
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

// Int8ToInt8 returns int8 int8_.
func Int8ToInt8(int8_ int8) int8 {
	return int8_
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
