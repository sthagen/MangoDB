// Code generated by "stringer -linecomment -type ErrorCode"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[errInternalError-1]
	_ = x[ErrBadValue-2]
	_ = x[ErrNamespaceNotFound-26]
	_ = x[ErrNamespaceExists-48]
	_ = x[ErrCommandNotFound-59]
	_ = x[ErrNotImplemented-238]
	_ = x[ErrRegexOptions-51075]
}

const (
	_ErrorCode_name_0 = "InternalErrorBadValue"
	_ErrorCode_name_1 = "NamespaceNotFound"
	_ErrorCode_name_2 = "NamespaceExists"
	_ErrorCode_name_3 = "CommandNotFound"
	_ErrorCode_name_4 = "NotImplemented"
	_ErrorCode_name_5 = "Location51075"
)

var (
	_ErrorCode_index_0 = [...]uint8{0, 13, 21}
)

func (i ErrorCode) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _ErrorCode_name_0[_ErrorCode_index_0[i]:_ErrorCode_index_0[i+1]]
	case i == 26:
		return _ErrorCode_name_1
	case i == 48:
		return _ErrorCode_name_2
	case i == 59:
		return _ErrorCode_name_3
	case i == 238:
		return _ErrorCode_name_4
	case i == 51075:
		return _ErrorCode_name_5
	default:
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
