package conv

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Slice conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// StringsToSlicePtrString parses string(s) strings_ to []*string.
func StringsToSlicePtrString(strings_ ...string) (slice_ptr_string_ []*string) {
	slice_ptr_string_ = make([]*string, len(strings_))

	for i, string_ := range strings_ {
		ptr_string_ := new(string)

		*ptr_string_ = string(string_)

		slice_ptr_string_[i] = ptr_string_
	}

	return
}

// StringsToSliceString parses string(s) strings_ to []string.
func StringsToSliceString(strings_ ...string) (slice_string_ []string) {
	return strings_
}

// StringsToPtrSlicePtrString parses string(s) strings_ to *[]*string.
func StringsToPtrSlicePtrString(strings_ ...string) (ptr_slice_ptr_string_ *[]*string) {
	slice_ptr_string_ := make([]*string, len(strings_))

	for i, string_ := range strings_ {
		ptr_string_ := new(string)

		*ptr_string_ = string(string_)

		slice_ptr_string_[i] = ptr_string_
	}

	ptr_slice_ptr_string_ = &slice_ptr_string_

	return
}

// StringsToPtrSliceString parses string(s) strings_ to *[]string.
func StringsToPtrSliceString(strings_ ...string) (ptr_slice_string_ *[]string) {
	slice_string_ := make([]string, len(strings_))

	copy(slice_string_, strings_)

	ptr_slice_string_ = &slice_string_

	return
}

// Uint64sToSlicePtrUint64 parses uint64(s) uint64s_ to []*uint64.
func Uint64sToSlicePtrUint64(uint64s_ ...uint64) (slice_ptr_uint64_ []*uint64) {
	slice_ptr_uint64_ = make([]*uint64, len(uint64s_))

	for i, uint64_ := range uint64s_ {
		ptr_uint64_ := new(uint64)

		*ptr_uint64_ = uint64(uint64_)

		slice_ptr_uint64_[i] = ptr_uint64_
	}

	return
}

// Uint64sToSliceUint64 parses uint64(s) uint64s_ to []uint64.
func Uint64sToSliceUint64(uint64s_ ...uint64) (slice_uint64_ []uint64) {
	return uint64s_
}

// Uint64sToPtrSlicePtrUint64 parses uint64(s) uint64s_ to *[]*uint64.
func Uint64sToPtrSlicePtrUint64(uint64s_ ...uint64) (ptr_slice_ptr_uint64_ *[]*uint64) {
	slice_ptr_uint64_ := make([]*uint64, len(uint64s_))

	for i, uint64_ := range uint64s_ {
		ptr_uint64_ := new(uint64)

		*ptr_uint64_ = uint64(uint64_)

		slice_ptr_uint64_[i] = ptr_uint64_
	}

	ptr_slice_ptr_uint64_ = &slice_ptr_uint64_

	return
}

// Uint64sToPtrSliceUint64 parses uint64(s) uint64s_ to *[]uint64.
func Uint64sToPtrSliceUint64(uint64s_ ...uint64) (ptr_slice_uint64_ *[]uint64) {
	slice_uint64_ := make([]uint64, len(uint64s_))

	copy(slice_uint64_, uint64s_)

	ptr_slice_uint64_ = &slice_uint64_

	return
}

// Uint32sToSlicePtrUint32 parses uint32(s) uint32s_ to []*uint32.
func Uint32sToSlicePtrUint32(uint32s_ ...uint32) (slice_ptr_uint32_ []*uint32) {
	slice_ptr_uint32_ = make([]*uint32, len(uint32s_))

	for i, uint32_ := range uint32s_ {
		ptr_uint32_ := new(uint32)

		*ptr_uint32_ = uint32(uint32_)

		slice_ptr_uint32_[i] = ptr_uint32_
	}

	return
}

// Uint32sToSliceUint32 parses uint32(s) uint32s_ to []uint32.
func Uint32sToSliceUint32(uint32s_ ...uint32) (slice_uint32_ []uint32) {
	return uint32s_
}

// Uint32sToPtrSlicePtrUint32 parses uint32(s) uint32s_ to *[]*uint32.
func Uint32sToPtrSlicePtrUint32(uint32s_ ...uint32) (ptr_slice_ptr_uint32_ *[]*uint32) {
	slice_ptr_uint32_ := make([]*uint32, len(uint32s_))

	for i, uint32_ := range uint32s_ {
		ptr_uint32_ := new(uint32)

		*ptr_uint32_ = uint32(uint32_)

		slice_ptr_uint32_[i] = ptr_uint32_
	}

	ptr_slice_ptr_uint32_ = &slice_ptr_uint32_

	return
}

// Uint32sToPtrSliceUint32 parses uint32(s) uint32s_ to *[]uint32.
func Uint32sToPtrSliceUint32(uint32s_ ...uint32) (ptr_slice_uint32_ *[]uint32) {
	slice_uint32_ := make([]uint32, len(uint32s_))

	copy(slice_uint32_, uint32s_)

	ptr_slice_uint32_ = &slice_uint32_

	return
}

// Uint16sToSlicePtrUint16 parses uint16(s) uint16s_ to []*uint16.
func Uint16sToSlicePtrUint16(uint16s_ ...uint16) (slice_ptr_uint16_ []*uint16) {
	slice_ptr_uint16_ = make([]*uint16, len(uint16s_))

	for i, uint16_ := range uint16s_ {
		ptr_uint16_ := new(uint16)

		*ptr_uint16_ = uint16(uint16_)

		slice_ptr_uint16_[i] = ptr_uint16_
	}

	return
}

// Uint16sToSliceUint16 parses uint16(s) uint16s_ to []uint16.
func Uint16sToSliceUint16(uint16s_ ...uint16) (slice_uint16_ []uint16) {
	return uint16s_
}

// Uint16sToPtrSlicePtrUint16 parses uint16(s) uint16s_ to *[]*uint16.
func Uint16sToPtrSlicePtrUint16(uint16s_ ...uint16) (ptr_slice_ptr_uint16_ *[]*uint16) {
	slice_ptr_uint16_ := make([]*uint16, len(uint16s_))

	for i, uint16_ := range uint16s_ {
		ptr_uint16_ := new(uint16)

		*ptr_uint16_ = uint16(uint16_)

		slice_ptr_uint16_[i] = ptr_uint16_
	}

	ptr_slice_ptr_uint16_ = &slice_ptr_uint16_

	return
}

// Uint16sToPtrSliceUint16 parses uint16(s) uint16s_ to *[]uint16.
func Uint16sToPtrSliceUint16(uint16s_ ...uint16) (ptr_slice_uint16_ *[]uint16) {
	slice_uint16_ := make([]uint16, len(uint16s_))

	copy(slice_uint16_, uint16s_)

	ptr_slice_uint16_ = &slice_uint16_

	return
}

// Uint8sToSlicePtrUint8 parses uint8(s) uint8s_ to []*uint8.
func Uint8sToSlicePtrUint8(uint8s_ ...uint8) (slice_ptr_uint8_ []*uint8) {
	slice_ptr_uint8_ = make([]*uint8, len(uint8s_))

	for i, uint8_ := range uint8s_ {
		ptr_uint8_ := new(uint8)

		*ptr_uint8_ = uint8(uint8_)

		slice_ptr_uint8_[i] = ptr_uint8_
	}

	return
}

// Uint8sToSliceUint8 parses uint8(s) uint8s_ to []uint8.
func Uint8sToSliceUint8(uint8s_ ...uint8) (slice_uint8_ []uint8) {
	return uint8s_
}

// Uint8sToPtrSlicePtrUint8 parses uint8(s) uint8s_ to *[]*uint8.
func Uint8sToPtrSlicePtrUint8(uint8s_ ...uint8) (ptr_slice_ptr_uint8_ *[]*uint8) {
	slice_ptr_uint8_ := make([]*uint8, len(uint8s_))

	for i, uint8_ := range uint8s_ {
		ptr_uint8_ := new(uint8)

		*ptr_uint8_ = uint8(uint8_)

		slice_ptr_uint8_[i] = ptr_uint8_
	}

	ptr_slice_ptr_uint8_ = &slice_ptr_uint8_

	return
}

// Uint8sToPtrSliceUint8 parses uint8(s) uint8s_ to *[]uint8.
func Uint8sToPtrSliceUint8(uint8s_ ...uint8) (ptr_slice_uint8_ *[]uint8) {
	slice_uint8_ := make([]uint8, len(uint8s_))

	copy(slice_uint8_, uint8s_)

	ptr_slice_uint8_ = &slice_uint8_

	return
}

// UintsToSlicePtrUint parses uint(s) uints_ to []*uint.
func UintsToSlicePtrUint(uints_ ...uint) (slice_ptr_uint_ []*uint) {
	slice_ptr_uint_ = make([]*uint, len(uints_))

	for i, uint_ := range uints_ {
		ptr_uint_ := new(uint)

		*ptr_uint_ = uint(uint_)

		slice_ptr_uint_[i] = ptr_uint_
	}

	return
}

// UintsToSliceUint parses uint(s) uints_ to []uint.
func UintsToSliceUint(uints_ ...uint) (slice_uint_ []uint) {
	return uints_
}

// UintsToPtrSlicePtrUint parses uint(s) uints_ to *[]*uint.
func UintsToPtrSlicePtrUint(uints_ ...uint) (ptr_slice_ptr_uint_ *[]*uint) {
	slice_ptr_uint_ := make([]*uint, len(uints_))

	for i, uint_ := range uints_ {
		ptr_uint_ := new(uint)

		*ptr_uint_ = uint(uint_)

		slice_ptr_uint_[i] = ptr_uint_
	}

	ptr_slice_ptr_uint_ = &slice_ptr_uint_

	return
}

// UintsToPtrSliceUint parses uint(s) uints_ to *[]uint.
func UintsToPtrSliceUint(uints_ ...uint) (ptr_slice_uint_ *[]uint) {
	slice_uint_ := make([]uint, len(uints_))

	copy(slice_uint_, uints_)

	ptr_slice_uint_ = &slice_uint_

	return
}

// Int64sToSlicePtrInt64 parses int64(s) int64s_ to []*int64.
func Int64sToSlicePtrInt64(int64s_ ...int64) (slice_ptr_int64_ []*int64) {
	slice_ptr_int64_ = make([]*int64, len(int64s_))

	for i, int64_ := range int64s_ {
		ptr_int64_ := new(int64)

		*ptr_int64_ = int64(int64_)

		slice_ptr_int64_[i] = ptr_int64_
	}

	return
}

// Int64sToSliceInt64 parses int64(s) int64s_ to []int64.
func Int64sToSliceInt64(int64s_ ...int64) (slice_int64_ []int64) {
	return int64s_
}

// Int64sToPtrSlicePtrInt64 parses int64(s) int64s_ to *[]*int64.
func Int64sToPtrSlicePtrInt64(int64s_ ...int64) (ptr_slice_ptr_int64_ *[]*int64) {
	slice_ptr_int64_ := make([]*int64, len(int64s_))

	for i, int64_ := range int64s_ {
		ptr_int64_ := new(int64)

		*ptr_int64_ = int64(int64_)

		slice_ptr_int64_[i] = ptr_int64_
	}

	ptr_slice_ptr_int64_ = &slice_ptr_int64_

	return
}

// Int64sToPtrSliceInt64 parses int64(s) int64s_ to *[]int64.
func Int64sToPtrSliceInt64(int64s_ ...int64) (ptr_slice_int64_ *[]int64) {
	slice_int64_ := make([]int64, len(int64s_))

	copy(slice_int64_, int64s_)

	ptr_slice_int64_ = &slice_int64_

	return
}

// Int32sToSlicePtrInt32 parses int32(s) int32s_ to []*int32.
func Int32sToSlicePtrInt32(int32s_ ...int32) (slice_ptr_int32_ []*int32) {
	slice_ptr_int32_ = make([]*int32, len(int32s_))

	for i, int32_ := range int32s_ {
		ptr_int32_ := new(int32)

		*ptr_int32_ = int32(int32_)

		slice_ptr_int32_[i] = ptr_int32_
	}

	return
}

// Int32sToSliceInt32 parses int32(s) int32s_ to []int32.
func Int32sToSliceInt32(int32s_ ...int32) (slice_int32_ []int32) {
	return int32s_
}

// Int32sToPtrSlicePtrInt32 parses int32(s) int32s_ to *[]*int32.
func Int32sToPtrSlicePtrInt32(int32s_ ...int32) (ptr_slice_ptr_int32_ *[]*int32) {
	slice_ptr_int32_ := make([]*int32, len(int32s_))

	for i, int32_ := range int32s_ {
		ptr_int32_ := new(int32)

		*ptr_int32_ = int32(int32_)

		slice_ptr_int32_[i] = ptr_int32_
	}

	ptr_slice_ptr_int32_ = &slice_ptr_int32_

	return
}

// Int32sToPtrSliceInt32 parses int32(s) int32s_ to *[]int32.
func Int32sToPtrSliceInt32(int32s_ ...int32) (ptr_slice_int32_ *[]int32) {
	slice_int32_ := make([]int32, len(int32s_))

	copy(slice_int32_, int32s_)

	ptr_slice_int32_ = &slice_int32_

	return
}

// Int16sToSlicePtrInt16 parses int16(s) int16s_ to []*int16.
func Int16sToSlicePtrInt16(int16s_ ...int16) (slice_ptr_int16_ []*int16) {
	slice_ptr_int16_ = make([]*int16, len(int16s_))

	for i, int16_ := range int16s_ {
		ptr_int16_ := new(int16)

		*ptr_int16_ = int16(int16_)

		slice_ptr_int16_[i] = ptr_int16_
	}

	return
}

// Int16sToSliceInt16 parses int16(s) int16s_ to []int16.
func Int16sToSliceInt16(int16s_ ...int16) (slice_int16_ []int16) {
	return int16s_
}

// Int16sToPtrSlicePtrInt16 parses int16(s) int16s_ to *[]*int16.
func Int16sToPtrSlicePtrInt16(int16s_ ...int16) (ptr_slice_ptr_int16_ *[]*int16) {
	slice_ptr_int16_ := make([]*int16, len(int16s_))

	for i, int16_ := range int16s_ {
		ptr_int16_ := new(int16)

		*ptr_int16_ = int16(int16_)

		slice_ptr_int16_[i] = ptr_int16_
	}

	ptr_slice_ptr_int16_ = &slice_ptr_int16_

	return
}

// Int16sToPtrSliceInt16 parses int16(s) int16s_ to *[]int16.
func Int16sToPtrSliceInt16(int16s_ ...int16) (ptr_slice_int16_ *[]int16) {
	slice_int16_ := make([]int16, len(int16s_))

	copy(slice_int16_, int16s_)

	ptr_slice_int16_ = &slice_int16_

	return
}

// Int8sToSlicePtrInt8 parses int8(s) int8s_ to []*int8.
func Int8sToSlicePtrInt8(int8s_ ...int8) (slice_ptr_int8_ []*int8) {
	slice_ptr_int8_ = make([]*int8, len(int8s_))

	for i, int8_ := range int8s_ {
		ptr_int8_ := new(int8)

		*ptr_int8_ = int8(int8_)

		slice_ptr_int8_[i] = ptr_int8_
	}

	return
}

// Int8sToSliceInt8 parses int8(s) int8s_ to []int8.
func Int8sToSliceInt8(int8s_ ...int8) (slice_int8_ []int8) {
	return int8s_
}

// Int8sToPtrSlicePtrInt8 parses int8(s) int8s_ to *[]*int8.
func Int8sToPtrSlicePtrInt8(int8s_ ...int8) (ptr_slice_ptr_int8_ *[]*int8) {
	slice_ptr_int8_ := make([]*int8, len(int8s_))

	for i, int8_ := range int8s_ {
		ptr_int8_ := new(int8)

		*ptr_int8_ = int8(int8_)

		slice_ptr_int8_[i] = ptr_int8_
	}

	ptr_slice_ptr_int8_ = &slice_ptr_int8_

	return
}

// Int8sToPtrSliceInt8 parses int8(s) int8s_ to *[]int8.
func Int8sToPtrSliceInt8(int8s_ ...int8) (ptr_slice_int8_ *[]int8) {
	slice_int8_ := make([]int8, len(int8s_))

	copy(slice_int8_, int8s_)

	ptr_slice_int8_ = &slice_int8_

	return
}

// IntsToSlicePtrInt parses int(s) ints_ to []*int.
func IntsToSlicePtrInt(ints_ ...int) (slice_ptr_int_ []*int) {
	slice_ptr_int_ = make([]*int, len(ints_))

	for i, int_ := range ints_ {
		ptr_int_ := new(int)

		*ptr_int_ = int(int_)

		slice_ptr_int_[i] = ptr_int_
	}

	return
}

// IntsToSliceInt parses int(s) ints_ to []int.
func IntsToSliceInt(ints_ ...int) (slice_int_ []int) {
	return ints_
}

// IntsToPtrSlicePtrInt parses int(s) ints_ to *[]*int.
func IntsToPtrSlicePtrInt(ints_ ...int) (ptr_slice_ptr_int_ *[]*int) {
	slice_ptr_int_ := make([]*int, len(ints_))

	for i, int_ := range ints_ {
		ptr_int_ := new(int)

		*ptr_int_ = int(int_)

		slice_ptr_int_[i] = ptr_int_
	}

	ptr_slice_ptr_int_ = &slice_ptr_int_

	return
}

// IntsToPtrSliceInt parses int(s) ints_ to *[]int.
func IntsToPtrSliceInt(ints_ ...int) (ptr_slice_int_ *[]int) {
	slice_int_ := make([]int, len(ints_))

	copy(slice_int_, ints_)

	ptr_slice_int_ = &slice_int_

	return
}

// BoolsToSlicePtrBool parses bool(s) bools_ to []*bool.
func BoolsToSlicePtrBool(bools_ ...bool) (slice_ptr_bool_ []*bool) {
	slice_ptr_bool_ = make([]*bool, len(bools_))

	for i, bool_ := range bools_ {
		ptr_bool_ := new(bool)

		*ptr_bool_ = bool(bool_)

		slice_ptr_bool_[i] = ptr_bool_
	}

	return
}

// BoolsToSliceBool parses bool(s) bools_ to []bool.
func BoolsToSliceBool(bools_ ...bool) (slice_bool_ []bool) {
	return bools_
}

// BoolsToPtrSlicePtrBool parses bool(s) bools_ to *[]*bool.
func BoolsToPtrSlicePtrBool(bools_ ...bool) (ptr_slice_ptr_bool_ *[]*bool) {
	slice_ptr_bool_ := make([]*bool, len(bools_))

	for i, bool_ := range bools_ {
		ptr_bool_ := new(bool)

		*ptr_bool_ = bool(bool_)

		slice_ptr_bool_[i] = ptr_bool_
	}

	ptr_slice_ptr_bool_ = &slice_ptr_bool_

	return
}

// BoolsToPtrSliceBool parses bool(s) bools_ to *[]bool.
func BoolsToPtrSliceBool(bools_ ...bool) (ptr_slice_bool_ *[]bool) {
	slice_bool_ := make([]bool, len(bools_))

	copy(slice_bool_, bools_)

	ptr_slice_bool_ = &slice_bool_

	return
}
