package string

import "strings"

type String struct{}

// StripEndingForwardSlash removes the ending forward slash of a string
func (s *String) StripEndingForwardSlash(value string) string {
	return StripEndingForwardSlash(value)
}

// StripLeading strips string of leading characters by an offset
func (s *String) StripLeading(value string, offset int) string {
	return StripLeading(value, offset)
}

// StripTrailing strips string of trailing characters after the length
func (s *String) StripTrailing(value string, length int) string {
	return StripTrailing(value, length)
}

// ToByte converts string or strings to byte array with an option for a separator
func (s *String) ToByte(sep string, elems ...string) []byte {
	return StringToByte(sep, elems...)
}

// StringToByte converts string or strings to byte array with an option for a separator
func StringToByte(sep string, elems ...string) []byte {
	return []byte(strings.Join(elems, sep))
}

// StripEndingForwardSlash removes the ending forward slash of a string
func StripEndingForwardSlash(value string) string {
	v := strings.TrimSpace(value)
	v = strings.ReplaceAll(v, `\`, `/`)
	ix := strings.LastIndex(v, `/`)
	if ix == (len(v) - 1) {
		return v[0:ix]
	}
	return v
}

// StripLeading strips string of leading characters by an offset
func StripLeading(value string, offset int) string {
	str := []rune(value)
	if len(str) > offset {
		return string(str[offset:])
	}
	return value
}

// StripTrailing strips string of trailing characters after the length
func StripTrailing(value string, length int) string {
	str := []rune(value)
	if len(str) > length {
		return string(str[0:length])
	}
	return value
}
