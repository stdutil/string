package string

import (
	"fmt"
	"strings"
)

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

// GetSides returns the start and end of the string value specified by length and join them together
//
// If the length of the value cannot leave a middle string, it will return the supplied value
func GetSides(value string, length int) string {
	if value == "" || length < 1 {
		return value
	}
	vlen := len(value)
	if length > vlen {
		return value
	}
	if length != 1 && vlen%length == 0 {
		return value
	}
	left := value[0:length]
	rght := value[vlen-length:]
	return left + rght
}

// NominateFileName returns a new file name based on the iteration
//
// If filename has no extension, the function will simply append the iteration number at the end
// If filename has an extension, it will insert the iteration number before the extension
func NominateFileName(origFN string, iteration int) string {
	extPos := strings.LastIndex(origFN, ".")
	if extPos == -1 {
		return fmt.Sprintf("%s [%d]", origFN, iteration)
	}

	ext := origFN[extPos:]
	bfn := origFN[:extPos]
	return fmt.Sprintf("%s_[%d]%s", bfn, iteration, ext)
}
