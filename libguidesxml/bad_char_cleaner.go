
package libguidesxml

import (
	"bufio"
	"unicode/utf8"
	
)

type BadCharCleaner struct {
	buffer *bufio.Reader
}

func (c BadCharCleaner) Read(b []byte) (n int, err error) {
	for {
		var r rune
		var s int
		r, s, err = c.buffer.ReadRune()
		if err != nil {
			return
		}
		if (r == '\u0001' || r == '\u0014' || r == '\u0019') && s == 1 {
			continue
		} else if n+s < len(b) {
			utf8.EncodeRune(b[n:], r)
			n += s
		} else {
			c.buffer.UnreadRune()
			break
		}
	}
	return
}