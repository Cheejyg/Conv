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

// StringToSlicePtrString parses a string string_ to []*string with provided separator.
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

// StringToSliceString parses a string string_ to []string with provided separator.
func StringToSliceString(string_, separator_ string) (slice_string_ []string) {
	slice_string_ = strings.Split(string_, separator_)

	return
}

// StringToPtrSlicePtrString parses a string string_ to *[]*string with provided separator.
func StringToPtrSlicePtrString(string_, separator_ string) (ptr_slice_ptr_string_ *[]*string) {
	slice_string_ := strings.Split(string_, separator_)

	slice_ptr_string_ := make([]*string, len(slice_string_))

	for i, string_ := range slice_string_ {
		ptr_string_ := new(string)

		*ptr_string_ = string(string_)

		slice_ptr_string_[i] = ptr_string_
	}

	ptr_slice_ptr_string_ = &slice_ptr_string_

	return
}

// StringToSliceString parses a string string_ to []string with provided separator.
func StringToPtrSliceString(string_, separator_ string) (ptr_slice_string_ *[]string) {
	slice_string_ := strings.Split(string_, separator_)

	ptr_slice_string_ = &slice_string_

	return
}
