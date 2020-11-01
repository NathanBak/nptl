package runes

import "unicode/utf8"

type Runes []rune

func FromBytes(buf []byte) (Runes, error) {
	return FromString(string(buf))
}

func FromString(s string) (Runes, error) {
	rs := []rune{}
	for _, r := range s {
		rs = append(rs, r)
	}
	return rs, nil
}

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
