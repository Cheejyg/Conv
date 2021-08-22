package Conv

import (
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import Conv "github.com/Cheejyg/Conv/Go"
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// String conversions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// StringToString returns string string_.
func StringToString(string_ string) string {
	return string_
}

// StringToPtr is equivalent to StringToPtrString(string_).
func StringToPtr(string_ string) (ptr_string_ *string) {
	ptr_string_ = StringToPtrString(string_)

	return
}

// StringToPtrString parses string string_ to *string.
func StringToPtrString(string_ string) (ptr_string_ *string) {
	ptr_string_ = new(string)

	*ptr_string_ = string(string_)

	return
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

	parse_bool_, err := strconv.ParseBool(string_)
	if err != nil {
		ptr_bool_ = nil

		return
	}

	*ptr_bool_ = bool(parse_bool_)

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
	parse_bool_, err := strconv.ParseBool(string_)
	if err != nil {
		return
	}

	bool_ = bool(parse_bool_)

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
