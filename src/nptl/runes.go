package nptl

import "unicode/utf8"

// Runes are used to pass around translatable content in an unambigous format.
type Runes []rune

// FromBytes reads a slice of bytes in ASCII or UTF-8 and then creates and returns new Runes.
func FromBytes(buf []byte) (Runes, error) {
	return FromString(string(buf))
}

// FromString creates and returns new Runes based on the provided string content.
func FromString(s string) (Runes, error) {
	rs := []rune{}
	for _, r := range s {
		rs = append(rs, r)
	}
	return rs, nil
}

// ToUTF8Bytes returns bytes representing the Runes in UTF-8 encoding.
func (rs Runes) ToUTF8Bytes() ([]byte, error) {
	buf := make([]byte, len(rs)*utf8.UTFMax)

	count := 0
	for _, r := range rs {
		count += utf8.EncodeRune(buf[count:], r)
	}
	buf = buf[:count]
	return buf, nil
}

func (rs Runes) String() string {
	return string(rs)
}
