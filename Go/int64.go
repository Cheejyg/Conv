package Conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import Conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int64 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int64ToString parses int64 int64_ to string.
func Int64ToString(int64_ int64) (string_ string) {
	string_ = fmt.Sprint(int64_)

	return
}

// Int64ToPtr is equivalent to Int64ToPtrInt64(int64_).
func Int64ToPtr(int64_ int64) (ptr_int64_ *int64) {
	ptr_int64_ = Int64ToPtrInt64(int64_)

	return
}

// Int64ToPtrString parses int64 int64_ to *string.
func Int64ToPtrString(int64_ int64) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int64_)

	return
}

// Int64ToPtrUint64 parses int64 int64_ to *uint64.
func Int64ToPtrUint64(int64_ int64) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int64_)

	return
}

// Int64ToPtrUint32 parses int64 int64_ to *uint32.
func Int64ToPtrUint32(int64_ int64) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int64_)

	return
}

// Int64ToPtrUint16 parses int64 int64_ to *uint16.
func Int64ToPtrUint16(int64_ int64) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int64_)

	return
}

// Int64ToPtrUint8 parses int64 int64_ to *uint8.
func Int64ToPtrUint8(int64_ int64) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int64_)

	return
}

// Int64ToPtrUint parses int64 int64_ to *uint.
func Int64ToPtrUint(int64_ int64) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int64_)

	return
}

// Int64ToPtrInt64 parses int64 int64_ to *int64.
func Int64ToPtrInt64(int64_ int64) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int64_)

	return
}

// Int64ToPtrInt32 parses int64 int64_ to *int32.
func Int64ToPtrInt32(int64_ int64) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int64_)

	return
}

// Int64ToPtrInt16 parses int64 int64_ to *int16.
func Int64ToPtrInt16(int64_ int64) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int64_)

	return
}

// Int64ToPtrInt8 parses int64 int64_ to *int8.
func Int64ToPtrInt8(int64_ int64) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int64_)

	return
}

// Int64ToPtrInt parses int64 int64_ to *int.
func Int64ToPtrInt(int64_ int64) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int64_)

	return
}

// Int64ToPtrBool parses int64 int64_ to *bool.
func Int64ToPtrBool(int64_ int64) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int64_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int64_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int64ToUint64 parses int64 int64_ to uint64.
func Int64ToUint64(int64_ int64) (uint64_ uint64) {
	uint64_ = uint64(int64_)

	return
}

// Int64ToUint32 parses int64 int64_ to uint32.
func Int64ToUint32(int64_ int64) (uint32_ uint32) {
	uint32_ = uint32(int64_)

	return
}

// Int64ToUint16 parses int64 int64_ to uint16.
func Int64ToUint16(int64_ int64) (uint16_ uint16) {
	uint16_ = uint16(int64_)

	return
}

// Int64ToUint8 parses int64 int64_ to uint8.
func Int64ToUint8(int64_ int64) (uint8_ uint8) {
	uint8_ = uint8(int64_)

	return
}

// Int64ToUint parses int64 int64_ to uint.
func Int64ToUint(int64_ int64) (uint_ uint) {
	uint_ = uint(int64_)

	return
}

// Int64ToInt64 returns int64 int64_.
func Int64ToInt64(int64_ int64) int64 {
	return int64_
}

// Int64ToInt32 parses int64 int64_ to int32.
func Int64ToInt32(int64_ int64) (int32_ int32) {
	int32_ = int32(int64_)

	return
}

// Int64ToInt16 parses int64 int64_ to int16.
func Int64ToInt16(int64_ int64) (int16_ int16) {
	int16_ = int16(int64_)

	return
}

// Int64ToInt8 parses int64 int64_ to int8.
func Int64ToInt8(int64_ int64) (int8_ int8) {
	int8_ = int8(int64_)

	return
}

// Int64ToInt parses int64 int64_ to int.
func Int64ToInt(int64_ int64) (int_ int) {
	int_ = int(int64_)

	return
}

// Int64ToBool parses int64 int64_ to bool.
func Int64ToBool(int64_ int64) (bool_ bool) {
	if int64_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int64_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
