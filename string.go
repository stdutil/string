package string

import "strings"

type String string

func (s *String) StripEndingForwardSlash() string {
	return StripEndingForwardSlash(string(*s))
}

func (s *String) StripLeading(offset int) string {
	return StripLeading(string(*s), offset)
}

func (s *String) StripTrailing(length int) string {
	return StripTrailing(string(*s), length)
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
