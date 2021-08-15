package conv

import (
	"fmt"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import conv "github.com/Cheejyg/conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ToString
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToString(interface_ interface{}) (string_ string) {
	return fmt.Sprint(interface_)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// strconv
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Atoi is equivalent to strconv.ParseInt(s, 10, 0), converted to type int.
func Atoi(string_ string) (int_ int, err error) {
	return strconv.Atoi(string_)
}

// Itoa is equivalent to strconv.FormatInt(int64(int_), 10).
func Itoa(int_ int) (string_ string) {
	return strconv.Itoa(int_)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// String conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// StringToString returns string string_.
func StringToString(string_ string) string {
	return string_
}

// StringToPtr is equivalent to StringToPtrString(string_).
func StringToPtr(string_ string) (ptr_string_ *string) {
	return StringToPtrString(string_)
}

// StringToPtrString parses string string_ to *string.
func StringToPtrString(string_ string) (ptr_string_ *string) {
	return &string_
}

// StringToPtrUint64 parses string string_ to *uint64.
func StringToPtrUint64(string_ string) (ptr_uint64_ *uint64) {
	ptr_uint64_ = new(uint64)

	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 64)
	if err != nil {
		ptr_uint64_ = nil

		return
	}

	*ptr_uint64_ = uint64(parse_uint_)

	return
}

// StringToPtrUint32 parses string string_ to *uint32.
func StringToPtrUint32(string_ string) (ptr_uint32_ *uint32) {
	ptr_uint32_ = new(uint32)

	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 32)
	if err != nil {
		ptr_uint32_ = nil

		return
	}

	*ptr_uint32_ = uint32(parse_uint_)

	return
}

// StringToPtrUint16 parses string string_ to *uint16.
func StringToPtrUint16(string_ string) (ptr_uint16_ *uint16) {
	ptr_uint16_ = new(uint16)

	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 16)
	if err != nil {
		ptr_uint16_ = nil

		return
	}

	*ptr_uint16_ = uint16(parse_uint_)

	return
}

// StringToPtrUint8 parses string string_ to *uint8.
func StringToPtrUint8(string_ string) (ptr_uint8_ *uint8) {
	ptr_uint8_ = new(uint8)

	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 8)
	if err != nil {
		ptr_uint8_ = nil

		return
	}

	*ptr_uint8_ = uint8(parse_uint_)

	return
}

// StringToPtrUint parses string string_ to *uint.
func StringToPtrUint(string_ string) (ptr_uint_ *uint) {
	ptr_uint_ = new(uint)

	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 0)
	if err != nil {
		ptr_uint_ = nil

		return
	}

	*ptr_uint_ = uint(parse_uint_)

	return
}

// StringToPtrInt64 parses string string_ to *int64.
func StringToPtrInt64(string_ string) (ptr_int64_ *int64) {
	ptr_int64_ = new(int64)

	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 64)
	if err != nil {
		ptr_int64_ = nil

		return
	}

	*ptr_int64_ = int64(parse_int_)

	return
}

// StringToPtrInt32 parses string string_ to *int32.
func StringToPtrInt32(string_ string) (ptr_int32_ *int32) {
	ptr_int32_ = new(int32)

	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 32)
	if err != nil {
		ptr_int32_ = nil

		return
	}

	*ptr_int32_ = int32(parse_int_)

	return
}

// StringToPtrInt16 parses string string_ to *int16.
func StringToPtrInt16(string_ string) (ptr_int16_ *int16) {
	ptr_int16_ = new(int16)

	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 16)
	if err != nil {
		ptr_int16_ = nil

		return
	}

	*ptr_int16_ = int16(parse_int_)

	return
}

// StringToPtrInt8 parses string string_ to *int8.
func StringToPtrInt8(string_ string) (ptr_int8_ *int8) {
	ptr_int8_ = new(int8)

	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 8)
	if err != nil {
		ptr_int8_ = nil

		return
	}

	*ptr_int8_ = int8(parse_int_)

	return
}

// StringToPtrInt parses string string_ to *int.
func StringToPtrInt(string_ string) (ptr_int_ *int) {
	ptr_int_ = new(int)

	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 0)
	if err != nil {
		ptr_int_ = nil

		return
	}

	*ptr_int_ = int(parse_int_)

	return
}

// StringToPtrBool parses string string_ to *bool.
func StringToPtrBool(string_ string) (ptr_bool_ *bool) {
	ptr_bool_ = new(bool)

	parse_bool, err := strconv.ParseBool(string_)
	if err != nil {
		ptr_bool_ = nil

		return
	}

	*ptr_bool_ = bool(parse_bool)

	/*switch strings.TrimSpace(string_) {
	case "0":
		fallthrough
	case "f":
		fallthrough
	case "F":
		fallthrough
	case "FALSE":
		fallthrough
	case "false":
		fallthrough
	case "False":
		*ptr_bool_ = false
	case "1":
		fallthrough
	case "t":
		fallthrough
	case "T":
		fallthrough
	case "TRUE":
		fallthrough
	case "true":
		fallthrough
	case "True":
		*ptr_bool_ = true
	default:
		ptr_bool_ = nil
	}*/

	return
}

// StringToUint64 parses string string_ to uint64.
func StringToUint64(string_ string) (uint64_ uint64) {
	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 64)
	if err != nil {
		uint64_ = ^uint64(0)

		return
	}

	uint64_ = uint64(parse_uint_)

	return
}

// StringToUint32 parses string string_ to uint32.
func StringToUint32(string_ string) (uint32_ uint32) {
	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 32)
	if err != nil {
		uint32_ = ^uint32(0)

		return
	}

	uint32_ = uint32(parse_uint_)

	return
}

// StringToUint16 parses string string_ to uint16.
func StringToUint16(string_ string) (uint16_ uint16) {
	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 16)
	if err != nil {
		uint16_ = ^uint16(0)

		return
	}

	uint16_ = uint16(parse_uint_)

	return
}

// StringToUint8 parses string string_ to uint8.
func StringToUint8(string_ string) (uint8_ uint8) {
	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 8)
	if err != nil {
		uint8_ = ^uint8(0)

		return
	}

	uint8_ = uint8(parse_uint_)

	return
}

// StringToUint parses string string_ to uint.
func StringToUint(string_ string) (uint_ uint) {
	parse_uint_, err := strconv.ParseUint(strings.TrimSpace(string_), 10, 0)
	if err != nil {
		uint_ = ^uint(0)

		return
	}

	uint_ = uint(parse_uint_)

	return
}

// StringToInt64 parses string string_ to int64.
func StringToInt64(string_ string) (int64_ int64) {
	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 64)
	if err != nil {
		int64_ = ^int64(0)

		return
	}

	int64_ = int64(parse_int_)

	return
}

// StringToInt32 parses string string_ to int32.
func StringToInt32(string_ string) (int32_ int32) {
	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 32)
	if err != nil {
		int32_ = ^int32(0)

		return
	}

	int32_ = int32(parse_int_)

	return
}

// StringToInt16 parses string string_ to int16.
func StringToInt16(string_ string) (int16_ int16) {
	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 16)
	if err != nil {
		int16_ = ^int16(0)

		return
	}

	int16_ = int16(parse_int_)

	return
}

// StringToInt8 parses string string_ to int8.
func StringToInt8(string_ string) (int8_ int8) {
	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 8)
	if err != nil {
		int8_ = ^int8(0)

		return
	}

	int8_ = int8(parse_int_)

	return
}

// StringToInt parses string string_ to int.
func StringToInt(string_ string) (int_ int) {
	parse_int_, err := strconv.ParseInt(strings.TrimSpace(string_), 10, 0)
	if err != nil {
		int_ = ^int(0)

		return
	}

	int_ = int(parse_int_)

	return
}

// StringToBool parses string string_ to bool.
func StringToBool(string_ string) (bool_ bool) {
	parse_bool, err := strconv.ParseBool(string_)
	if err != nil {
		return
	}

	bool_ = bool(parse_bool)

	/* switch strings.TrimSpace(string_) {
	case "0":
		fallthrough
	case "f":
		fallthrough
	case "F":
		fallthrough
	case "FALSE":
		fallthrough
	case "false":
		fallthrough
	case "False":
		bool_ = false
	case "1":
		fallthrough
	case "t":
		fallthrough
	case "T":
		fallthrough
	case "TRUE":
		fallthrough
	case "true":
		fallthrough
	case "True":
		bool_ = true
	default:
	} */

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Slice conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
	return &uint64s_
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
	return &uint32s_
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
	return &uint16s_
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
	return &uint8s_
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
	return &uints_
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
	return &int64s_
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
	return &int32s_
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
	return &int16s_
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
	return &int8s_
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
	return &ints_
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
	return &bools_
}
