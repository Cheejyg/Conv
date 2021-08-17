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
	string_ = fmt.Sprint(interface_)

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// strconv
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Atoi is equivalent to strconv.ParseInt(s, 10, 0), converted to type int.
func Atoi(string_ string) (int_ int, err error) {
	int_, err = strconv.Atoi(string_)

	return
}

// Itoa is equivalent to strconv.FormatInt(int64(int_), 10).
func Itoa(int_ int) (string_ string) {
	string_ = strconv.Itoa(int_)

	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// String conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// StringToSlicePtrString parses string string_ to []*string with provided separator.
func StringToSlicePtrString(string_, separator_ string) (slice_ptr_string_ []*string) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_string_ = make([]*string, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_string_ := new(string)

		*ptr_string_ = string(string_)

		slice_ptr_string_[i] = ptr_string_
	}

	return
}

// StringToSliceString parses string string_ to []string with provided separator.
func StringToSliceString(string_, separator_ string) (slice_string_ []string) {
	slice_string_ = strings.Split(string_, separator_)

	return
}

// StringToPtrSlicePtrString parses string string_ to *[]*string with provided separator.
func StringToPtrSlicePtrString(string_, separator_ string) (ptr_slice_ptr_string_ *[]*string) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_string_ := make([]*string, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_string_ := StringToPtr(string_)

		slice_ptr_string_[i] = ptr_string_
	}

	ptr_slice_ptr_string_ = &slice_ptr_string_

	return
}

// StringToPtrSliceString parses string string_ to *[]string with provided separator.
func StringToPtrSliceString(string_, separator_ string) (ptr_slice_string_ *[]string) {
	slice_string_ := strings.Split(string_, separator_)

	ptr_slice_string_ = &slice_string_

	return
}

// StringToSlicePtrUint64 parses string string_ to []*uint64 with provided separator.
func StringToSlicePtrUint64(string_, separator_ string) (slice_ptr_uint64_ []*uint64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint64_ = make([]*uint64, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint64_ := new(uint64)

		ptr_uint64_ = StringToPtrUint64(string_)
		if ptr_uint64_ == nil {
			slice_ptr_uint64_ = nil

			return
		}

		slice_ptr_uint64_[i] = ptr_uint64_
	}

	return
}

// StringToSliceUint64 parses string string_ to []uint64 with provided separator.
func StringToSliceUint64(string_, separator_ string) (slice_uint64_ []uint64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint64_ = make([]uint64, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint64_ := StringToUint64(string_)

		slice_uint64_[i] = uint64_
	}

	return
}

// StringToPtrSlicePtrUint64 parses string string_ to *[]*uint64 with provided separator.
func StringToPtrSlicePtrUint64(string_, separator_ string) (ptr_slice_ptr_uint64_ *[]*uint64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint64_ := make([]*uint64, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint64_ := StringToPtrUint64(string_)
		if ptr_uint64_ == nil {
			ptr_slice_ptr_uint64_ = nil

			return
		}

		slice_ptr_uint64_[i] = ptr_uint64_
	}

	ptr_slice_ptr_uint64_ = &slice_ptr_uint64_

	return
}

// StringToPtrSliceUint64 parses string string_ to *[]uint64 with provided separator.
func StringToPtrSliceUint64(string_, separator_ string) (ptr_slice_uint64_ *[]uint64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint64_ := make([]uint64, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint64_ := StringToUint64(string_)

		slice_uint64_[i] = uint64_
	}

	ptr_slice_uint64_ = &slice_uint64_

	return
}

// StringToSlicePtrUint32 parses string string_ to []*uint32 with provided separator.
func StringToSlicePtrUint32(string_, separator_ string) (slice_ptr_uint32_ []*uint32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint32_ = make([]*uint32, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint32_ := new(uint32)

		ptr_uint32_ = StringToPtrUint32(string_)
		if ptr_uint32_ == nil {
			slice_ptr_uint32_ = nil

			return
		}

		slice_ptr_uint32_[i] = ptr_uint32_
	}

	return
}

// StringToSliceUint32 parses string string_ to []uint32 with provided separator.
func StringToSliceUint32(string_, separator_ string) (slice_uint32_ []uint32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint32_ = make([]uint32, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint32_ := StringToUint32(string_)

		slice_uint32_[i] = uint32_
	}

	return
}

// StringToPtrSlicePtrUint32 parses string string_ to *[]*uint32 with provided separator.
func StringToPtrSlicePtrUint32(string_, separator_ string) (ptr_slice_ptr_uint32_ *[]*uint32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint32_ := make([]*uint32, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint32_ := StringToPtrUint32(string_)
		if ptr_uint32_ == nil {
			ptr_slice_ptr_uint32_ = nil

			return
		}

		slice_ptr_uint32_[i] = ptr_uint32_
	}

	ptr_slice_ptr_uint32_ = &slice_ptr_uint32_

	return
}

// StringToPtrSliceUint32 parses string string_ to *[]uint32 with provided separator.
func StringToPtrSliceUint32(string_, separator_ string) (ptr_slice_uint32_ *[]uint32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint32_ := make([]uint32, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint32_ := StringToUint32(string_)

		slice_uint32_[i] = uint32_
	}

	ptr_slice_uint32_ = &slice_uint32_

	return
}

// StringToSlicePtrUint16 parses string string_ to []*uint16 with provided separator.
func StringToSlicePtrUint16(string_, separator_ string) (slice_ptr_uint16_ []*uint16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint16_ = make([]*uint16, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint16_ := new(uint16)

		ptr_uint16_ = StringToPtrUint16(string_)
		if ptr_uint16_ == nil {
			slice_ptr_uint16_ = nil

			return
		}

		slice_ptr_uint16_[i] = ptr_uint16_
	}

	return
}

// StringToSliceUint16 parses string string_ to []uint16 with provided separator.
func StringToSliceUint16(string_, separator_ string) (slice_uint16_ []uint16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint16_ = make([]uint16, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint16_ := StringToUint16(string_)

		slice_uint16_[i] = uint16_
	}

	return
}

// StringToPtrSlicePtrUint16 parses string string_ to *[]*uint16 with provided separator.
func StringToPtrSlicePtrUint16(string_, separator_ string) (ptr_slice_ptr_uint16_ *[]*uint16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint16_ := make([]*uint16, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint16_ := StringToPtrUint16(string_)
		if ptr_uint16_ == nil {
			ptr_slice_ptr_uint16_ = nil

			return
		}

		slice_ptr_uint16_[i] = ptr_uint16_
	}

	ptr_slice_ptr_uint16_ = &slice_ptr_uint16_

	return
}

// StringToPtrSliceUint16 parses string string_ to *[]uint16 with provided separator.
func StringToPtrSliceUint16(string_, separator_ string) (ptr_slice_uint16_ *[]uint16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint16_ := make([]uint16, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint16_ := StringToUint16(string_)

		slice_uint16_[i] = uint16_
	}

	ptr_slice_uint16_ = &slice_uint16_

	return
}

// StringToSlicePtrUint8 parses string string_ to []*uint8 with provided separator.
func StringToSlicePtrUint8(string_, separator_ string) (slice_ptr_uint8_ []*uint8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint8_ = make([]*uint8, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint8_ := new(uint8)

		ptr_uint8_ = StringToPtrUint8(string_)
		if ptr_uint8_ == nil {
			slice_ptr_uint8_ = nil

			return
		}

		slice_ptr_uint8_[i] = ptr_uint8_
	}

	return
}

// StringToSliceUint8 parses string string_ to []uint8 with provided separator.
func StringToSliceUint8(string_, separator_ string) (slice_uint8_ []uint8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint8_ = make([]uint8, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint8_ := StringToUint8(string_)

		slice_uint8_[i] = uint8_
	}

	return
}

// StringToPtrSlicePtrUint8 parses string string_ to *[]*uint8 with provided separator.
func StringToPtrSlicePtrUint8(string_, separator_ string) (ptr_slice_ptr_uint8_ *[]*uint8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint8_ := make([]*uint8, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint8_ := StringToPtrUint8(string_)
		if ptr_uint8_ == nil {
			ptr_slice_ptr_uint8_ = nil

			return
		}

		slice_ptr_uint8_[i] = ptr_uint8_
	}

	ptr_slice_ptr_uint8_ = &slice_ptr_uint8_

	return
}

// StringToPtrSliceUint8 parses string string_ to *[]uint8 with provided separator.
func StringToPtrSliceUint8(string_, separator_ string) (ptr_slice_uint8_ *[]uint8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint8_ := make([]uint8, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint8_ := StringToUint8(string_)

		slice_uint8_[i] = uint8_
	}

	ptr_slice_uint8_ = &slice_uint8_

	return
}

// StringToSlicePtrUint parses string string_ to []*uint with provided separator.
func StringToSlicePtrUint(string_, separator_ string) (slice_ptr_uint_ []*uint) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint_ = make([]*uint, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint_ := new(uint)

		ptr_uint_ = StringToPtrUint(string_)
		if ptr_uint_ == nil {
			slice_ptr_uint_ = nil

			return
		}

		slice_ptr_uint_[i] = ptr_uint_
	}

	return
}

// StringToSliceUint parses string string_ to []uint with provided separator.
func StringToSliceUint(string_, separator_ string) (slice_uint_ []uint) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint_ = make([]uint, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint_ := StringToUint(string_)

		slice_uint_[i] = uint_
	}

	return
}

// StringToPtrSlicePtrUint parses string string_ to *[]*uint with provided separator.
func StringToPtrSlicePtrUint(string_, separator_ string) (ptr_slice_ptr_uint_ *[]*uint) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_uint_ := make([]*uint, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_uint_ := StringToPtrUint(string_)
		if ptr_uint_ == nil {
			ptr_slice_ptr_uint_ = nil

			return
		}

		slice_ptr_uint_[i] = ptr_uint_
	}

	ptr_slice_ptr_uint_ = &slice_ptr_uint_

	return
}

// StringToPtrSliceUint parses string string_ to *[]uint with provided separator.
func StringToPtrSliceUint(string_, separator_ string) (ptr_slice_uint_ *[]uint) {
	slice_string_ := strings.Split(string_, separator_)

	slice_uint_ := make([]uint, len(slice_string_))

	for i, string_ := range slice_string_ {
		uint_ := StringToUint(string_)

		slice_uint_[i] = uint_
	}

	ptr_slice_uint_ = &slice_uint_

	return
}

// StringToSlicePtrInt64 parses string string_ to []*int64 with provided separator.
func StringToSlicePtrInt64(string_, separator_ string) (slice_ptr_int64_ []*int64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int64_ = make([]*int64, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int64_ := new(int64)

		ptr_int64_ = StringToPtrInt64(string_)
		if ptr_int64_ == nil {
			slice_ptr_int64_ = nil

			return
		}

		slice_ptr_int64_[i] = ptr_int64_
	}

	return
}

// StringToSliceInt64 parses string string_ to []int64 with provided separator.
func StringToSliceInt64(string_, separator_ string) (slice_int64_ []int64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int64_ = make([]int64, len(slice_string_))

	for i, string_ := range slice_string_ {
		int64_ := StringToInt64(string_)

		slice_int64_[i] = int64_
	}

	return
}

// StringToPtrSlicePtrInt64 parses string string_ to *[]*int64 with provided separator.
func StringToPtrSlicePtrInt64(string_, separator_ string) (ptr_slice_ptr_int64_ *[]*int64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int64_ := make([]*int64, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int64_ := StringToPtrInt64(string_)
		if ptr_int64_ == nil {
			ptr_slice_ptr_int64_ = nil

			return
		}

		slice_ptr_int64_[i] = ptr_int64_
	}

	ptr_slice_ptr_int64_ = &slice_ptr_int64_

	return
}

// StringToPtrSliceInt64 parses string string_ to *[]int64 with provided separator.
func StringToPtrSliceInt64(string_, separator_ string) (ptr_slice_int64_ *[]int64) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int64_ := make([]int64, len(slice_string_))

	for i, string_ := range slice_string_ {
		int64_ := StringToInt64(string_)

		slice_int64_[i] = int64_
	}

	ptr_slice_int64_ = &slice_int64_

	return
}

// StringToSlicePtrInt32 parses string string_ to []*int32 with provided separator.
func StringToSlicePtrInt32(string_, separator_ string) (slice_ptr_int32_ []*int32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int32_ = make([]*int32, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int32_ := new(int32)

		ptr_int32_ = StringToPtrInt32(string_)
		if ptr_int32_ == nil {
			slice_ptr_int32_ = nil

			return
		}

		slice_ptr_int32_[i] = ptr_int32_
	}

	return
}

// StringToSliceInt32 parses string string_ to []int32 with provided separator.
func StringToSliceInt32(string_, separator_ string) (slice_int32_ []int32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int32_ = make([]int32, len(slice_string_))

	for i, string_ := range slice_string_ {
		int32_ := StringToInt32(string_)

		slice_int32_[i] = int32_
	}

	return
}

// StringToPtrSlicePtrInt32 parses string string_ to *[]*int32 with provided separator.
func StringToPtrSlicePtrInt32(string_, separator_ string) (ptr_slice_ptr_int32_ *[]*int32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int32_ := make([]*int32, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int32_ := StringToPtrInt32(string_)
		if ptr_int32_ == nil {
			ptr_slice_ptr_int32_ = nil

			return
		}

		slice_ptr_int32_[i] = ptr_int32_
	}

	ptr_slice_ptr_int32_ = &slice_ptr_int32_

	return
}

// StringToPtrSliceInt32 parses string string_ to *[]int32 with provided separator.
func StringToPtrSliceInt32(string_, separator_ string) (ptr_slice_int32_ *[]int32) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int32_ := make([]int32, len(slice_string_))

	for i, string_ := range slice_string_ {
		int32_ := StringToInt32(string_)

		slice_int32_[i] = int32_
	}

	ptr_slice_int32_ = &slice_int32_

	return
}

// StringToSlicePtrInt16 parses string string_ to []*int16 with provided separator.
func StringToSlicePtrInt16(string_, separator_ string) (slice_ptr_int16_ []*int16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int16_ = make([]*int16, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int16_ := new(int16)

		ptr_int16_ = StringToPtrInt16(string_)
		if ptr_int16_ == nil {
			slice_ptr_int16_ = nil

			return
		}

		slice_ptr_int16_[i] = ptr_int16_
	}

	return
}

// StringToSliceInt16 parses string string_ to []int16 with provided separator.
func StringToSliceInt16(string_, separator_ string) (slice_int16_ []int16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int16_ = make([]int16, len(slice_string_))

	for i, string_ := range slice_string_ {
		int16_ := StringToInt16(string_)

		slice_int16_[i] = int16_
	}

	return
}

// StringToPtrSlicePtrInt16 parses string string_ to *[]*int16 with provided separator.
func StringToPtrSlicePtrInt16(string_, separator_ string) (ptr_slice_ptr_int16_ *[]*int16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int16_ := make([]*int16, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int16_ := StringToPtrInt16(string_)
		if ptr_int16_ == nil {
			ptr_slice_ptr_int16_ = nil

			return
		}

		slice_ptr_int16_[i] = ptr_int16_
	}

	ptr_slice_ptr_int16_ = &slice_ptr_int16_

	return
}

// StringToPtrSliceInt16 parses string string_ to *[]int16 with provided separator.
func StringToPtrSliceInt16(string_, separator_ string) (ptr_slice_int16_ *[]int16) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int16_ := make([]int16, len(slice_string_))

	for i, string_ := range slice_string_ {
		int16_ := StringToInt16(string_)

		slice_int16_[i] = int16_
	}

	ptr_slice_int16_ = &slice_int16_

	return
}

// StringToSlicePtrInt8 parses string string_ to []*int8 with provided separator.
func StringToSlicePtrInt8(string_, separator_ string) (slice_ptr_int8_ []*int8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int8_ = make([]*int8, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int8_ := new(int8)

		ptr_int8_ = StringToPtrInt8(string_)
		if ptr_int8_ == nil {
			slice_ptr_int8_ = nil

			return
		}

		slice_ptr_int8_[i] = ptr_int8_
	}

	return
}

// StringToSliceInt8 parses string string_ to []int8 with provided separator.
func StringToSliceInt8(string_, separator_ string) (slice_int8_ []int8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int8_ = make([]int8, len(slice_string_))

	for i, string_ := range slice_string_ {
		int8_ := StringToInt8(string_)

		slice_int8_[i] = int8_
	}

	return
}

// StringToPtrSlicePtrInt8 parses string string_ to *[]*int8 with provided separator.
func StringToPtrSlicePtrInt8(string_, separator_ string) (ptr_slice_ptr_int8_ *[]*int8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int8_ := make([]*int8, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int8_ := StringToPtrInt8(string_)
		if ptr_int8_ == nil {
			ptr_slice_ptr_int8_ = nil

			return
		}

		slice_ptr_int8_[i] = ptr_int8_
	}

	ptr_slice_ptr_int8_ = &slice_ptr_int8_

	return
}

// StringToPtrSliceInt8 parses string string_ to *[]int8 with provided separator.
func StringToPtrSliceInt8(string_, separator_ string) (ptr_slice_int8_ *[]int8) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int8_ := make([]int8, len(slice_string_))

	for i, string_ := range slice_string_ {
		int8_ := StringToInt8(string_)

		slice_int8_[i] = int8_
	}

	ptr_slice_int8_ = &slice_int8_

	return
}

// StringToSlicePtrInt parses string string_ to []*int with provided separator.
func StringToSlicePtrInt(string_, separator_ string) (slice_ptr_int_ []*int) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int_ = make([]*int, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int_ := new(int)

		ptr_int_ = StringToPtrInt(string_)
		if ptr_int_ == nil {
			slice_ptr_int_ = nil

			return
		}

		slice_ptr_int_[i] = ptr_int_
	}

	return
}

// StringToSliceInt parses string string_ to []int with provided separator.
func StringToSliceInt(string_, separator_ string) (slice_int_ []int) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int_ = make([]int, len(slice_string_))

	for i, string_ := range slice_string_ {
		int_ := StringToInt(string_)

		slice_int_[i] = int_
	}

	return
}

// StringToPtrSlicePtrInt parses string string_ to *[]*int with provided separator.
func StringToPtrSlicePtrInt(string_, separator_ string) (ptr_slice_ptr_int_ *[]*int) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_int_ := make([]*int, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_int_ := StringToPtrInt(string_)
		if ptr_int_ == nil {
			ptr_slice_ptr_int_ = nil

			return
		}

		slice_ptr_int_[i] = ptr_int_
	}

	ptr_slice_ptr_int_ = &slice_ptr_int_

	return
}

// StringToPtrSliceInt parses string string_ to *[]int with provided separator.
func StringToPtrSliceInt(string_, separator_ string) (ptr_slice_int_ *[]int) {
	slice_string_ := strings.Split(string_, separator_)

	slice_int_ := make([]int, len(slice_string_))

	for i, string_ := range slice_string_ {
		int_ := StringToInt(string_)

		slice_int_[i] = int_
	}

	ptr_slice_int_ = &slice_int_

	return
}

// StringToSlicePtrBool parses string string_ to []*bool with provided separator.
func StringToSlicePtrBool(string_, separator_ string) (slice_ptr_bool_ []*bool) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_bool_ = make([]*bool, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_bool_ := new(bool)

		ptr_bool_ = StringToPtrBool(string_)
		if ptr_bool_ == nil {
			slice_ptr_bool_ = nil

			return
		}

		slice_ptr_bool_[i] = ptr_bool_
	}

	return
}

// StringToSliceBool parses string string_ to []bool with provided separator.
func StringToSliceBool(string_, separator_ string) (slice_bool_ []bool) {
	slice_string_ := strings.Split(string_, separator_)

	slice_bool_ = make([]bool, len(slice_string_))

	for i, string_ := range slice_string_ {
		bool_ := StringToBool(string_)

		slice_bool_[i] = bool_
	}

	return
}

// StringToPtrSlicePtrBool parses string string_ to *[]*bool with provided separator.
func StringToPtrSlicePtrBool(string_, separator_ string) (ptr_slice_ptr_bool_ *[]*bool) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_bool_ := make([]*bool, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_bool_ := StringToPtrBool(string_)
		if ptr_bool_ == nil {
			ptr_slice_ptr_bool_ = nil

			return
		}

		slice_ptr_bool_[i] = ptr_bool_
	}

	ptr_slice_ptr_bool_ = &slice_ptr_bool_

	return
}

// StringToPtrSliceBool parses string string_ to *[]bool with provided separator.
func StringToPtrSliceBool(string_, separator_ string) (ptr_slice_bool_ *[]bool) {
	slice_string_ := strings.Split(string_, separator_)

	slice_bool_ := make([]bool, len(slice_string_))

	for i, string_ := range slice_string_ {
		bool_ := StringToBool(string_)

		slice_bool_[i] = bool_
	}

	ptr_slice_bool_ = &slice_bool_

	return
}
