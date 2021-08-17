package conv

import "fmt"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Int16 conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Int16ToString parses int16 int16_ to string.
func Int16ToString(int16_ int16) (string_ string) {
	string_ = fmt.Sprint(int16_)

	return
}

// Int16ToPtr is equivalent to Int16ToPtrInt16(int16_).
func Int16ToPtr(int16_ int16) (ptr_int16_ *int16) {
	ptr_int16_ = Int16ToPtrInt16(int16_)

	return
}

// Int16ToPtrString parses int16 int16_ to *string.
func Int16ToPtrString(int16_ int16) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = fmt.Sprint(int16_)

	return
}

// Int16ToPtrUint64 parses int16 int16_ to *uint64.
func Int16ToPtrUint64(int16_ int16) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	*ptr_uint64_ = uint64(int16_)

	return
}

// Int16ToPtrUint32 parses int16 int16_ to *uint32.
func Int16ToPtrUint32(int16_ int16) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	*ptr_uint32_ = uint32(int16_)

	return
}

// Int16ToPtrUint16 parses int16 int16_ to *uint16.
func Int16ToPtrUint16(int16_ int16) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	*ptr_uint16_ = uint16(int16_)

	return
}

// Int16ToPtrUint8 parses int16 int16_ to *uint8.
func Int16ToPtrUint8(int16_ int16) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	*ptr_uint8_ = uint8(int16_)

	return
}

// Int16ToPtrUint parses int16 int16_ to *uint.
func Int16ToPtrUint(int16_ int16) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	*ptr_uint_ = uint(int16_)

	return
}

// Int16ToPtrInt64 parses int16 int16_ to *int64.
func Int16ToPtrInt64(int16_ int16) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	*ptr_int64_ = int64(int16_)

	return
}

// Int16ToPtrInt32 parses int16 int16_ to *int32.
func Int16ToPtrInt32(int16_ int16) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	*ptr_int32_ = int32(int16_)

	return
}

// Int16ToPtrInt16 parses int16 int16_ to *int16.
func Int16ToPtrInt16(int16_ int16) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	*ptr_int16_ = int16(int16_)

	return
}

// Int16ToPtrInt8 parses int16 int16_ to *int8.
func Int16ToPtrInt8(int16_ int16) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	*ptr_int8_ = int8(int16_)

	return
}

// Int16ToPtrInt parses int16 int16_ to *int.
func Int16ToPtrInt(int16_ int16) (ptr_int_ *int) {
	ptr_int_ = new(int)

	*ptr_int_ = int(int16_)

	return
}

// Int16ToPtrBool parses int16 int16_ to *bool.
func Int16ToPtrBool(int16_ int16) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	if int16_ == 0 { // || int_ < 0 {
		*ptr_bool_ = false
	} else if int16_ == 1 { // || int_ > 0 {
		*ptr_bool_ = true
	} else {
		ptr_bool_ = nil
	}

	return
}

// Int16ToUint64 parses int16 int16_ to uint64.
func Int16ToUint64(int16_ int16) (uint64_ uint64) {
	uint64_ = uint64(int16_)

	return
}

// Int16ToUint32 parses int16 int16_ to uint32.
func Int16ToUint32(int16_ int16) (uint32_ uint32) {
	uint32_ = uint32(int16_)

	return
}

// Int16ToUint16 parses int16 int16_ to uint16.
func Int16ToUint16(int16_ int16) (uint16_ uint16) {
	uint16_ = uint16(int16_)

	return
}

// Int16ToUint8 parses int16 int16_ to uint8.
func Int16ToUint8(int16_ int16) (uint8_ uint8) {
	uint8_ = uint8(int16_)

	return
}

// Int16ToUint parses int16 int16_ to uint.
func Int16ToUint(int16_ int16) (uint_ uint) {
	uint_ = uint(int16_)

	return
}

// Int16ToInt64 parses int16 int16_ to int64.
func Int16ToInt64(int16_ int16) (int64_ int64) {
	int64_ = int64(int16_)

	return
}

// Int16ToInt32 parses int16 int16_ to int32.
func Int16ToInt32(int16_ int16) (int32_ int32) {
	int32_ = int32(int16_)

	return
}

// Int16ToInt16 returns int16 int16_.
func Int16ToInt16(int16_ int16) int16 {
	return int16_
}

// Int16ToInt8 parses int16 int16_ to int8.
func Int16ToInt8(int16_ int16) (int8_ int8) {
	int8_ = int8(int16_)

	return
}

// Int16ToInt parses int16 int16_ to int.
func Int16ToInt(int16_ int16) (int_ int) {
	int_ = int(int16_)

	return
}

// Int16ToBool parses int16 int16_ to bool.
func Int16ToBool(int16_ int16) (bool_ bool) {
	if int16_ == 0 { // || int_ < 0 {
		bool_ = false
	} else if int16_ == 1 { // || int_ > 0 {
		bool_ = true
	} else {
	}

	return
}
