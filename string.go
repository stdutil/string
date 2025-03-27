package string

import (
	"fmt"
	"strconv"
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

// NominateFileName returns a new file name by appending a number
func NominateFileName(origFN string) string {
	var (
		bfn, ext string
		num      int64
		err      error
	)

	bfn = origFN

	// Check for a file extension
	extPos := strings.LastIndex(origFN, ".")
	if extPos != -1 && extPos > 0 {
		bfn = origFN[:extPos]
		ext = origFN[extPos:]
	}

	// Check for a number inside a parenthesis
	// This will indicate that this is a redundancy of the file
	if prp := strings.LastIndex(bfn, ")"); prp != -1 {
		sblp := bfn[0:prp]
		if plp := strings.LastIndex(sblp, "("); plp != -1 {
			if rem := sblp[plp+1:]; rem != "" {
				num, err = strconv.ParseInt(rem, 10, 64)
				if err == nil {
					bfn = sblp[0:plp]
				}
			}
		}
	}
	bfn = strings.TrimSpace(bfn)
	return fmt.Sprintf("%s (%d)%s", bfn, num+1, ext)
}
